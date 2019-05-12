// init delete file modal
(function($){
    $(document).ready(function(){

        var deleteFileModal = $("#fe26-delete-file-modal").data("fe26Modal");

        deleteFileModal.onOk = function () {

            let data = {
                filepath: $('#fe26-delete-file-modal input[name="filename"]')[0].value,
                action: "delete-file",
            };

            $.post( "fe26.json",data, function( data ) {
                setTimeout(function(){ location.reload() }, 250);
            });
        };

        deleteFileModal.afterOpen = function () {
            $("#fe26-delete-file-modal .fe26-modal-content")[0].tabIndex=-1;
            $("#fe26-delete-file-modal .fe26-modal-content")[0].focus();
        };


        // deleteFileModal.afterClose = function () {
        //     console.log("afterClose");
        // };

        $(".fe26-delete-file").click(function (e) {
            e.preventDefault();

            var filePath = $(this).attr("path");
            var filename = $(this).parent().parent().find(".fe26-col-name a").html()

            $('#fe26-delete-file-modal input[name="filename"]')[0].value = filePath;
            $('#fe26-delete-file-modal .filename').html(filename);
            deleteFileModal.open();
        });

        $(".fe26-c2c-file").each(function () {
            $(this).click(function (e) {
                e.preventDefault();
                let port = window.location.port;
                if (port != "" && port != 0){
                    port =":"+port;
                }
                let url = window.location.protocol+"//"+window.location.hostname+port + $(this).attr("href");
                $.fn.c2c(url);
                fe26Dialog().info("Link copied to clipboard");
            });
        });


    });
})(jQuery);