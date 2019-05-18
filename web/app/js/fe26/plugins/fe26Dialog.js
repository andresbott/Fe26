window.fe26Dialog = function(opts){

    var _fe26Dialog = function(opts){

        var defaults = {
            timeout:1500,
            infoIcon:""
        };
        var options = $.extend(defaults, opts);
        var dialogContainer = $('<div class="fe26-dialog-container"></div>');

        var constructor = function (){
            $("body").prepend(dialogContainer);
        };
        constructor();

        this.info = function (txt) {
            doDialog("info",txt);
        };

        var dialogRemove = function(dialog){
            dialog.animate({
                opacity: 0,
                bottom: "+=15",
            }, 300, function() {
                this.remove();
            });
        };

        var dialogAdd =  function (dialog) {
            dialogContainer.append(dialog);
            dialog.css({opacity:0}).animate({
                opacity: 1,
                bottom: "+=15",
            }, 300);
        };

        var doDialog = function (type,msg){
            var t ="";
            var icon ="";
            if (type == "info"){
                t = "fe26-dialog-info";
                if (options.infoIcon != ""){
                    icon = options.infoIcon;
                }
            }


            var dialog = $('<div class="fe26-dialog '+t+'">'+icon+msg+'</div>');
            dialogAdd(dialog);

            setTimeout(function () {
                dialogRemove(dialog);
            },options.timeout);
        }

    };

    // Return early if this element already has a plugin instance
    if ($(window).data('fe26Dialog')) return $(window).data('fe26Dialog')
    var thisFe26Dialog = new _fe26Dialog(opts);
    $(window).data('fe26Dialog', thisFe26Dialog);
    return $(window).data('fe26Dialog');
};