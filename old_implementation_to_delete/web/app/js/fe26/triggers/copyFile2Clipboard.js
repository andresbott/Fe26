// copy to clipboard
$(".fe26-c2c-file").each(function () {

    fe26Dialog({
        infoIcon:'<i class=\"fe26i-ok fe26-dialog-icon\"></i>',
        // timeout:5000000
    });

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
