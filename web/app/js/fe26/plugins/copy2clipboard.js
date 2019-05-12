(function($){
    $.fn.c2c = function(txt){
        var $temp = $("<textarea>");
        $("body").append($temp);
        $temp.val(txt).select();
        document.execCommand("copy");
        $temp.remove();
    }
})(jQuery);
