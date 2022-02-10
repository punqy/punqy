let area = $('<div id="draggable-area"></div>');
const body = $('body');
const Document = $(document);
(() => {
    setTimeout(() => {
        // ============ VARS START ============//
        const header = $(".info");
        header.empty();
        const swaggerData = JSON.parse($("#swagger-data").html());
        let autoAuthMode = localStorage.getItem('aa') === '1';
        let authorizeButton = $(".authorize");
        let defaultAuth = swaggerData.spec.security[0].default_auth;
        let areas = swaggerData.spec.security[0].areas;
        let clientCredentials = JSON.parse(defaultAuth.client_credentials);
        let userCredentials = JSON.parse(defaultAuth.user_credentials.replace(/\'/g, ''));
        let switchUserSelect = $(`<select></select>`);
        let optBlockSections = $(".opblock-tag-section");
        let searchInput = $('<input type=\"search\" placeholder=\"Search tag\" id=\"search\">');

        // ============ VARS END ============//
        const menu = [];

        menu.push(new MenuButton({
            text: swaggerData.spec.info.title,
            attributes: {
                class: 'codename',
                onclick: 'return false;'
            },
        }));

        // ============ AREAS START ============//
        if (typeof areas !== 'undefined' && areas.constructor.name === 'Array') {
            areas.unshift('');
            // $(".info").find("h2").append('<span style="float: right" id="areas-links"></span>');

            areas.forEach(a => {
                const name = a !== '' ? a : 'default';
                const disabled = window.location.pathname === ('/api/doc/' + a).replace(/\/$/, "");
                const item = new MenuButton({
                    text: name[0].toUpperCase() + name.slice(1),
                    attributes: {
                        href: `/api/doc/${a}`,
                        class: `area-link ${disabled ? 'disabled' : ''}`,
                        onclick: `${disabled ? 'return false;' : ''}`
                    },
                });
                menu.push(item);
            });
        }
        menu.push(new MenuButton({
            text: '',
            attributes: {
                class: 'space areas'
            },
        }));

        menu.push(new MenuButton({
            attributes: {
                class: `auto-auth-btn ${localStorage.getItem('aa') === '1' ? 'auth-enabled' : ''}`,
            },
            handlers: {
                click: (e, item) => {
                    if (localStorage.getItem('aa') === '1') {
                        item.removeClass('auth-enabled');
                        return localStorage.removeItem('aa');
                    }
                    item.addClass('auth-enabled');
                    localStorage.setItem('aa', '1');
                }
            },
        }));
        // ============ AREAS END ============//

        // ============ SEARCH START ============//
        $(".auth-wrapper").prepend(searchInput);

        searchInput.keyup(c => {
            let b = $(c.target).val();
            optBlockSections.each(function (a, c) {
                return $(c).find(".nostyle").find("span").text().toLowerCase().match(new RegExp(b.toLowerCase())) ? void $(c).show() : void $(c).hide()
            });
        });

        // ============ SEARCH END ============//

        // ============ FAST SCROLL START ============//

        let toTopButton = $("<a id=\"toTop\"><span>&#9650</span></a>");
        let toBottomButton = $("<a id=\"toBottom\"><span>&#9660</span></a>");

        toTopButton.on("click", function (a) {
            a.preventDefault();
            $("html, body").animate({scrollTop: 0}, "300")
        });

        toBottomButton.on("click", function (a) {
            a.preventDefault();
            $("html, body").animate({scrollTop: Document.height()}, "300")
        });
        body.append(toTopButton).append(toBottomButton);

        $(window).scroll(function () {
            300 < $(window).scrollTop() ? toTopButton.show() : toTopButton.hide();
            300 < $(window).scrollTop() ? toBottomButton.hide() : toBottomButton.show()
        });
        // ============ FAST SCROLL END ============//

        // ============ AUTH START ============//

        userCredentials.forEach((e, i) => {
            let option = $(`<option value="${e.password}">${e.username}</option>`);
            option.prop('selected', e.username === localStorage.getItem('current-username'));
            switchUserSelect.append(option);
        });
        menu.push(new MenuButton({
            text: 'Authorize default',
            attributes: {
                class: 'authorize_default'
            },
            handlers: {
                click: (e, item) => {
                    let user = switchUserSelect.find('option:selected');
                    let authData = {...clientCredentials, username: user.text(), password: user.val()};
                    token = '';
                    $.post("/api/v1/oauth/token", authData).done(a => {
                        token = a.payload.access_token
                    })
                }
            },
            init: (item) => {
                if (autoAuthMode) {
                    item.click();
                }
            }
        }));
        switchUserSelect.change(e => {
            if (switchUserSelect.val() === '') {
                ascPass()
            }
            let user = switchUserSelect.find('option:selected');
            console.log(user);
            localStorage.setItem('current-username', user.text());
            $('.authorize_default').click();
        });

        function ascPass() {
            let modal = new tingle.modal({
                footer: true,
                stickyFooter: false,
                closeMethods: ['overlay', 'button', 'escape'],
                closeLabel: "Close",
                cssClass: ['custom-class-1', 'custom-class-2'],
                onOpen: function () {
                },
                onClose: function () {
                },
                beforeClose: function () {
                    return true; // close the modal
                }
            });

            modal.setContent('<div class="swagger-ui"><h3 >Enter password</h3><input class="" type="password" placeholder="password"></div>');
            modal.addFooterBtn('Submit', 'tingle-btn tingle-btn--primary tingle-btn--pull-right', function () {
                modal.close();
            });
            modal.open();
        }

        document.onkeyup = function (a) {
            a.ctrlKey && a.altKey && 65 === a.which && $('.authorize_default').trigger("click")
        };


        Document.on('click', '#auto-auth', function (e) {

            if (localStorage.getItem('aa')) {
                $(e.target).css('background', '#7d8492').html('Enable auto auth &#x2713;');
                return localStorage.removeItem('aa');
            }
            $(e.target).css('background', '#ee821b').html('Disable auto auth X');
            localStorage.setItem('aa', 1);
        });

        // ============ AUTH END ============//
        menu.push(new MenuButton({
            text: 'Documentation',
            attributes: {
                target: '_blank',
                id: 'documentation-btn',
                href: 'https://drive.google.com/drive/folders/1Kxlf6cJa3yWTPvys5pElKKAHlMFPr_Es',
                class: 'documentation-link'
            },
        }));

        menu.forEach(item => {
            header.append(item.print());
        });


        authorizeButton
            .after(switchUserSelect);

    }, 0);

    // ============ DATE_PICKERS START ============//

    // Init datepickers
    body.on('click', '.try-out__btn', (e) => {
        let target = $(e.target);
        let section = target.closest('.opblock-section');
        let tbody = section.find('tbody');

        setTimeout(() => {
            let trs = tbody.find('tr');
            trs.each((i, e) => {
                let tr = $(e);
                if (/date/.test(tr.find('.prop-format').text())) {
                    let input = $(tr.find('input'));
                    input.datepicker({
                        dateFormat: "yy-mm-dd",
                        onClose: function (date) {
                            let ev2 = new Event('input', {bubbles: true});
                            ev2.simulated = true;
                            input[0].value = date;
                            input[0].dispatchEvent(ev2);
                        }
                    });
                }
            });

        }, 10);
    });
    // ============ DATE_PICKERS END ============//

    // ============ WINDOWS START ============//

    Document.on('click', '.close-frame', function (event) {
        $(event.target).closest('.iframe-container').remove();
    });
    body.append(area);
    Document.on('click', '.debug-link', function (event, a) {
        event.preventDefault();
        let target = $(event.target);
        let className = target.text();
        let url = target.prop('href');

        let cWindow = new Window(className, url, area);
        let instance = cWindow.open();

        target.prop('target', 'frame' + className);
        let frame = instance.find('iframe');
        frame.load(function (e) {
            frame.contents().find("#header").remove();
            frame.contents().find("body").append($("<style type='text/css'> #sidebar,#sidebar-shortcuts{background-color: #52555F} </style>"));
            frame.contents().find("#sidebar-shortcuts").css({"background-color": " #52555F"});
        });
    });

    // ============ WINDOWS END ============//

})();


class MenuButton {
    constructor(options) {
        this.wrapper = $('<span></span>');
        this.button = $('<span></span>');

        if (options.attributes) {
            Object.keys(options.attributes).forEach(a => {
                this.button.prop(a, options.attributes[a]);
            });
        }

        if (options.handlers) {
            Object.keys(options.handlers).forEach(evt => {
                this.button.on(evt, (e) => {
                    options.handlers[evt](e, this.button)
                });
            });
        }
        this.button.html(this.buildInner(options));
        this.button.addClass('swagger-menu-item');
        this.wrapper.append(this.button);
        if (options.init) {
            options.init(this.button);
        }
    }

    buildInner(options) {
        const link = options.text || '';
        if (options.attributes && options.attributes.href) {
            return $(`<a onclick="${options.attributes.onclick || ''}" href="${options.attributes.href}" target="${options.attributes.target || '_self'}">${options.text}</a>`);
        }
        return $(`<a href="#">${options.text || ''}</a>`);
    }

    print() {
        return this.wrapper;
    }
}


/**
 * class window
 */
class Window {

    /**
     * @param className
     * @param src
     * @param area
     * @param callback
     */
    constructor(className, src, area, callback) {
        this.className = className;
        this.src = src;
        this.area = area;
        this.callback = callback;
    }

    /**
     * @return {jQuery.fn.init|jQuery|HTMLElement}
     */
    open() {

        let self = this;

        let cWindow = $(`<div class="window" style="position:fixed"></div>`);
        let title = $(`<span class="window-title" >${self.className}</span>`);
        let minimizeButton = $(`<a class="minimize"><i class="fa fa-window-minimize text-light" style="font-size: 12px"></i></a>`);
        let closeButton = $(`<a class="close-window"><i class="fa fa-times text-light"></i></a>`);
        let newTabButton = $(`<a href="${self.src}" target="_blank" class="new-tab"><i class="fa fa-window-maximize text-light" style="font-size: 12px"></i></a>`);
        let frame = $(`<iframe src="${self.src}" frameborder="0" name="frame${self.className}" class="frame"></iframe>`);

        /**
         * Add unique name
         */
        $(cWindow, title, minimizeButton, closeButton, newTabButton, frame).addClass(self.className);

        /**
         * Add nodes
         */
        cWindow.append([title, newTabButton, minimizeButton, closeButton, frame]);

        /**
         * Close window when new tab opening
         */
        $(newTabButton).click(() => {
            closeButton.click();
        });

        /**
         * Make window draggable
         */
        cWindow.draggable({
            iframeFix: true,
            containment: self.area,
            start: function () {
                $('.window').css('z-index', 1);
                cWindow.css('z-index', 2);
                let x = window.scrollX;
                let y = window.scrollY;
                window.onscroll = () => {
                    window.scrollTo(x, y);
                };
            },
            stop: function () {
                window.onscroll = () => {
                    window.scrollTo(0, 0);
                };
                window.onscroll = () => {
                };
            }
        });

        /**
         * Make window minimizable
         */
        minimizeButton.click(() => {
            if (cWindow.hasClass('mini')) {
                return cWindow.removeClass('mini');
            }
            cWindow.addClass('mini')
        });

        /**
         * Make window closable
         */
        closeButton.click(() => {
            cWindow.remove();
        });

        /**
         * Push window
         */
        body.append(cWindow);

        /**
         * Call callback
         */
        if (typeof this.callback === 'function') {
            self.callback(cWindow, frame, title, closeButton, newTabButton, minimizeButton);
        }

        return cWindow;
    };
}