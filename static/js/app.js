'use strict';

(function () {

    let current_token = "";

    /**
     * 初始化
     */
    let init = function () {
        refreshMessageList();
        bindMessageForm();
        bindRemoveBtn();
        refreshUIElementStatus();
        bindLoginForm();
    };

    let refreshUIElementStatus = function () {
        if (current_token) {
            $('[data-auth-required="logging"]').show();
            $('[data-auth-required="unlogining"]').hide();
        } else {
            $('[data-auth-required="logging"]').hide();
            $('[data-auth-required="unlogining"]').show();
        }
    };


    /**
     * 刷新留言列表
     */
    let refreshMessageList = function () {
        request('/message/list', {}, function (data) {
            let html = '';
            let html_tpl = $('#template-message-box').html()
            for (let i in data.list) {
                if (!data.list.hasOwnProperty(i)) continue;
                let item_html = html_tpl;
                for (let v in data.list[i]) {
                    if (!data.list[i].hasOwnProperty(v)) continue;
                    item_html = item_html.replaceAll("{{" + v + "}}", data.list[i][v])
                }
                html += item_html;
            }
            $('#message-list').html(html);

            bindRemoveBtn();
        });
    };

    /**
     * 绑定留言表单
     */
    let bindMessageForm = function () {
        $('#new-message-submit').unbind().bind('click', function () {
            if (confirm('确定要提交留言吗？')) {
                let payload = {
                    author: $("#new-message-author").val(),
                    title: $("#new-message-title").val(),
                    content: $("#new-message-content").val(),
                };
                request('/message/create', payload, function () {
                    refreshMessageList();
                })
            }
        });
    };

    let bindLoginForm = function () {
        $('#login-submit').unbind().bind('click', function () {
            if (confirm('确定要提交留言吗？')) {
                let payload = {
                    username: $("#login-username").val(),
                    password: $("#login-password").val(),
                };
                request('/auth/login', payload, function (data) {
                    current_token = data['token'];
                    refreshUIElementStatus();
                    refreshMessageList();
                })
            }
        });
    };

    /**
     * 绑定删除按钮
     */
    let bindRemoveBtn = function () {
        $('button[data-func="remove-message"]').unbind().bind('click', function () {
            let id = parseInt($(this).attr('data-id'));
            if (id > 0) {
                if (confirm('确定要删除留言吗？')) {
                    request('/message/remove', {id: id}, function () {
                        refreshMessageList();
                    });
                }
            }
        });
    }

    /**
     * 封装 API
     * @param api
     * @param payload
     * @param success_callback
     * @param fail_callback
     */
    let request = function (api, payload, success_callback, fail_callback) {
        // API请求带上TOKEN
        payload['_token'] = current_token;
        // 请求API
        $.post('/api' + api, JSON.stringify(payload), function (ret, http_status) {
            if (ret.hasOwnProperty('code') && ret.hasOwnProperty('data')) {
                if (ret.code === 'success') {
                    success_callback(ret.data)
                    return null;
                }
            }
            onFail(ret)
        }).fail(
            function (fail_object) {
                let content = fail_object.responseText;
                let content_obj = JSON.parse(content);
                if (content_obj) {
                    onFail(content_obj)
                } else {
                    onFail(content)
                }
            }
        );
        // 失败处理
        let onFail = function (ret) {
            console.log('[API FAIL]', api, payload, ret)
            if (ret.hasOwnProperty('code')) {
                if (ret.code === 'auth_required') {
                    current_token = '';
                    refreshUIElementStatus();
                }
            }

            if (ret.hasOwnProperty('message')) {
                alert('[操作失败] ' + ret.message);
            }
            if (fail_callback) {
                fail_callback(ret)
            }
        }
    }

    init();
})()