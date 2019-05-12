'use strict';
(function($){
    window.fe26Dialog = function(){



        var _fe26Dialog = function(options){

            var defaults = {
                dialogTimeout:1500,
            };
            var options = $.extend(defaults, options);
            var dialogContainer = $('<div class="fe26-dialog-container"></div>');

            var constructor = function (){
                $("body").append(dialogContainer);
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



                // dialog.hide(function () {

                //     this.animate({
                //         opacity: 1,
                //       }, 300, function() {
                //
                //     });
                // })
            };

            var doDialog = function (type,msg){
                var t ="";
                if (type == "info"){
                    t = "fe26-dialog-info";
                }

                var dialog = $('<div class="fe26-dialog '+t+'">'+msg+'</div>');
                dialogAdd(dialog);

                setTimeout(function () {
                    dialogRemove(dialog);
                },options.dialogTimeout);
            }

        };

        // Return early if this element already has a plugin instance
        if ($(window).data('fe26Dialog')) return $(window).data('fe26Dialog')
        var thisFe26Dialog = new _fe26Dialog();
        $(window).data('fe26Dialog', thisFe26Dialog);
        return $(window).data('fe26Dialog');
    }
})(jQuery);

window.fe26Dialog();