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