// init Create Dir modal


var newDirModal = $("#fe26-create-dir-modal").data("fe26Modal");

newDirModal.onOk = function () {
    let data = {
        path: $('#fe26-create-dir-modal input[name="path"]')[0].value,
        foldername: $('#fe26-create-dir-modal input[name="foldername"]')[0].value,
        action: "create-dir",
    };
    $.post( "fe26.json",data, function( data ) {
        console.log(data)
        setTimeout(function(){ location.reload() }, 250);
    });
};

newDirModal.afterOpen = function () {
    $('#fe26-create-dir-modal input[name="foldername"]')[0].focus();
};

newDirModal.afterClose = function () {
    $('#fe26-create-dir-modal input[name="foldername"]')[0].value = "";
};

$("#fe26-create-folder-btn").click(function (e) {
    e.preventDefault();
    newDirModal.open();
});

