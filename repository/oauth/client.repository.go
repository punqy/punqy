package oauth

import (
	"context"
	"fmt"
	punqy "github.com/punqy/core"
	model "github.com/punqy/punqy/model/storage"
	"github.com/punqy/punqy/repository/tables"
	"github.com/slmder/qbuilder"
	"reflect"
	"strings"
)

type ClientRepository interface {
	Find(ctx context.Context, id string) (punqy.OAuthClient, error)
	FindOneByClientIdSecretAndGrantType(ctx context.Context, cID, sec string, gt punqy.GrantType) (punqy.OAuthClient, error)
	Insert(ctx context.Context, entity *model.OAuthClient) error
	NewOauthClient(ctx context.Context) (model.OAuthClient, error)
}

type clientRepository struct {
	punqy.Dal
}

func NewClientRepository(db punqy.Dal) ClientRepository {
	return &clientRepository{Dal: db}
}

func (r *clientRepository) NewOauthClient(ctx context.Context) (model.OAuthClient, error) {
	e := model.OAuthClient{
		GrantTypes:   model.StringList{punqy.GrantTypeRefreshToken.String(), punqy.ClientCredentials.String(), punqy.GrantTypePassword.String()},
		ClientSecret: fmt.Sprintf("cs_%s", punqy.RandomString(64)),
	}
	if err := e.NewId(); err != nil {
		return e, err
	}

	if err := r.Insert(ctx, &e); err != nil {
		return model.OAuthClient{}, err
	}
	return e, nil
}

func (r *clientRepository) Insert(ctx context.Context, entity *model.OAuthClient) error {
	sql := r.BuildInsert(tables.OAuthClient).
		Row(punqy.StringMap{
			"id":                  ":id",
			"allowed_grant_types": ":allowed_grant_types",
			"client_secret":       ":client_secret",
			"created_at":          "now()",
			"updated_at":          "now()",
		}).ToSQL()

	_, err := r.DoInsert(ctx, sql, entity)
	return r.PipeErr(err)
}

func (r *clientRepository) FindOneByClientIdSecretAndGrantType(ctx context.Context, cID, sec string, gt punqy.GrantType) (punqy.OAuthClient, error) {
	var entity model.OAuthClient
	query := r.BuildSelectE(entity).
		From(tables.OAuthClient).
		AndWhere("id = $1").
		AndWhere("client_secret = $2").
		AndWhere("$3 IN (select jsonb_array_elements_text(allowed_grant_types))").
		Limit(1).
		ToSQL()
	println(FieldList(entity))
	return entity, r.PipeErr(r.DoSelectOne(ctx, &entity, query, cID, sec, gt))
}

func (r *clientRepository) FindBy(ctx context.Context, cond qbuilder.Conditions, pager punqy.Pagination) ([]model.OAuthClient, error) {
	var entities []model.OAuthClient
	return entities, r.Dal.FindBy(ctx, tables.OAuthClient, &entities, cond, pager)
}

func (r *clientRepository) FindOneBy(ctx context.Context, cond qbuilder.Conditions) (model.OAuthClient, error) {
	entity := model.OAuthClient{}
	return entity, r.Dal.FindOneBy(ctx, tables.OAuthClient, &entity, cond)
}

func (r *clientRepository) Find(ctx context.Context, id string) (punqy.OAuthClient, error) {
	return r.FindOneBy(ctx, qbuilder.Conditions{"id": id})
}

func FieldList(obj interface{}, alias ...string) string {
	objType := reflect.TypeOf(obj)
	switch objType.Kind() {
	case reflect.Ptr:
		val := objType.Elem()
		if val.Kind() != reflect.Struct {
			argErr := fmt.Errorf("source must be a struct or struct pointer %s given", val.Kind().String())
			panic(argErr)
		}
		objType = val
	case reflect.Struct:
	default:
		argErr := fmt.Errorf("source must be a struct or struct pointer %s given", objType.Kind().String())
		panic(argErr)
	}
	names := make([]string, 0)
	println(objType.NumField())
	for i := 0; i < objType.NumField(); i++ {
		field := objType.Field(i)
		if db, ok := field.Tag.Lookup("db"); ok {
			if db != "" {
				name := db
				if len(alias) > 0 && alias[0] != "" {
					name = fmt.Sprintf("%s.%s", alias[0], db)
				}
				names = append(names, name)
			}
		}
	}
	return strings.Join(names, ", ")
}

func Names(fd reflect.StructField, names *[]string, alias ...string) {
	if db, ok := fd.Tag.Lookup("db"); ok {
		if db != "" {
			name := db
			if len(alias) > 0 && alias[0] != "" {
				name = fmt.Sprintf("%s.%s", alias[0], db)
			}
			*names = append(*names, name)
		}
	}
	println(reflect.TypeOf(fd).Kind().String())
	if !fd.Anonymous && reflect.TypeOf(fd).Kind() == reflect.Struct  {
		for i := 0; i < fd.Type.NumField(); i++ {
			field := fd.Type.Field(i)
			Names(field, names,alias...)
		}
	}
}

func parseName(field reflect.StructField, tagName string) (tag, fieldName string) {
	// first, set the fieldName to the field's name
	fieldName = field.Name
	// if a mapFunc is set, use that to override the fieldName

	// if there's no tag to look for, return the field name
	if tagName == "" {
		return "", fieldName
	}

	// if this tag is not set using the normal convention in the tag,
	// then return the fieldname..  this check is done because according
	// to the reflect documentation:
	//    If the tag does not have the conventional format,
	//    the value returned by Get is unspecified.
	// which doesn't sound great.
	if !strings.Contains(string(field.Tag), tagName+":") {
		return "", fieldName
	}

	// at this point we're fairly sure that we have a tag, so lets pull it out
	tag = field.Tag.Get(tagName)


	// finally, split the options from the name
	parts := strings.Split(tag, ",")
	fieldName = parts[0]

	return tag, fieldName
}
