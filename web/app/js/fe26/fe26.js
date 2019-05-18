'use strict';
// in an ideal world I would make this as an react spa
// include jquery
@@include('../../vendor/jquery/jquery-3.4.1.min.js')
(function($){
@@include('./plugins/copy2clipboard.js')
@@include('./plugins/fe26Dialog.js')
@@include('./plugins/fe26Modal.js')
    $(document).ready(function(){
@@include('./plugins/fe26FileUpload.js')
@@include('./triggers/createDirModal.js')
@@include('./triggers/deleteFileModal.js')
@@include('./triggers/copyFile2Clipboard.js')
    });
})(jQuery);