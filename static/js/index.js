//import "common"

$('.dropdown').click(function () {
    $('.dropdown-menu').toggle('show');
});

$('.formAjax').submit(function () {
    var _this = this;
    var method = $(this).attr('method');
    var action = $(this).attr('action');

    if (action.indexOf('?') === -1) {
        action = action + '?_t=' + Math.random()
    } else {
        action = action + '&_t=' + Math.random()
    }

    $.ajax({
        type: method.toLowerCase(),
        url: action,
        data: $(this).serialize(),  // 序列化表单数据
        cache: false,
        dataType: 'json',
        success: function (data) {
            if (data.func) {
                eval(data.func + "()");
                return false;
            }

            if (data.code == 1) {
                var json = data.data;
                if (json.url) {
                    layer.msg(data.msg, {
                        icon: 1,
                        time: 2000 //2秒关闭（如果不配置，默认是3秒）
                    }, function(){
                        window.location.href = json.url;
                    });
                    return false;
                }

                layer.alert(data.msg, {icon: 1});
            }

            if (data.code == 2) {
                layer.alert(data.msg, {icon: 2});
            }
        },
        error: function () {
            layer.alert('网络错误，请查看网络', {icon: 2});
        }
    });

    return false;   // 禁止默认 submit 事件
});

