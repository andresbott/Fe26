
$.fn.fe26Modal = function(options){

    var _fe26Modal = function(element){
        var modal = $(element); // the container div
        var okBtn = null;
        var cancelBtn = null;

        // var closeBtn = null;
        var obj = this;  // this object set in a variable to be accesed from within other functions
        // var defaults = {
        //     mouseOverStopSlide:true,
        // };

        var transitionEvent = function (){
            var t;
            var el = document.createElement('fakeelement');
            var transitions = {
                'transition':'transitionend',
                'OTransition':'oTransitionEnd',
                'MozTransition':'transitionend',
                'WebkitTransition':'webkitTransitionEnd'
            };

            for(t in transitions){
                if( el.style[t] !== undefined ){
                    return transitions[t];
                }
            }
        }();

        var keyboardHandler = function (e) {
            if(e.key === "Escape") {
                obj.close();
            }
            if(e.key === "Enter") {
                obj.ok();
            }
        };

        this.open = function () {
            if ( ! modal.hasClass("fe26-modal-visible")){
                modal.addClass("fe26-modal-visible");
                $(window).keydown(keyboardHandler);

                if(typeof(obj.afterOpen) == "function" ){
                    var control = true;
                    $(modal[0]).bind(transitionEvent,function () {
                        if (control){
                            obj.afterOpen();
                            control = false;
                        }
                    })
                }
            }
        };

        this.ok = function () {
            if(typeof(obj.onOk()) == "function" ){ obj.onOk();  }
            obj.close();
        };

        this.close = function () {
            if ( modal.hasClass("fe26-modal-visible")) {
                modal.removeClass("fe26-modal-visible");
                $(window).unbind('keydown');
                if(typeof(obj.afterClose()) == "function" ){ obj.afterClose();  }
            }
        };

        this.afterClose = function () {
            // dummy function to extend functionality
        };

        this.onOk = function () {
            // dummy function to extend functionality
        };

        this.afterOpen = function () {
            // dummy function to extend functionality
        };

        var constructor = function (){
            let $modal = modal.detach();
            $("body").prepend($modal);

            let $content = modal.find(".fe26-modal-content").detach();
            let valign = $('<div class="fe26-modal-valign"></div>');
            let actions = $('<div class="fe26-modal-actions"></div>');

            okBtn = $('<div class="fe26-modal-button fe26-button-ok">' +
                '<span class="glyphicon glyphicon-ok" aria-hidden="true"></span><span>OK</span>' +
                '</div>');
            okBtn.click(function (e) {
                e.preventDefault();
                obj.ok();
            });
            actions.append(okBtn);


            cancelBtn = $('<div class="fe26-modal-button fe26-button-cancel">' +
                '<span class="glyphicon glyphicon-remove" aria-hidden="true"></span><span>Cancel</span>' +
                '</div>');
            cancelBtn.click(function (e) {
                e.preventDefault();
                obj.close();
            });

            actions.append(cancelBtn);
            $content.append(actions);
            valign.append($content);
            modal.append(valign)
        };
        constructor();
    };

    return this.each(function() {
        var element = $(this);
        // Return early if this element already has a plugin instance
        if (element.data('fe26Modal')) return;
        var thisFe26Modal = new _fe26Modal(this,options);
        element.data('fe26Modal', thisFe26Modal);
    });
};

$(document).ready(function(){
    $(".fe26-modal").fe26Modal();
});
