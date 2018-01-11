
$(document).ready(function () {
    productID = parseInt($("#product_ID").val());

    $("#product_updateform").submit( function(e) {
        e.preventDefault();
        var data = {};
        data.product_id = -1
        data.product_isbn = $("#product_isbn").val()
        data.product_publisher = $("#product_publisher").val()
        data.product_author = $("#product_author").val()
        data.product_genre = $("#product_genre").val()
        data.product_description = $("#product_description").val()
        data.product_name = $("#product_name").val()
        data.product_type = $("#product_type").val()
        


        $.ajax({
            url: "/v1/product",
            dataType: 'json',
            type: 'POST',
            contentType: 'application/json',
            data: JSON.stringify(data),
            success: function (data) {
                window.location.href="/product/get/"+data.id
            },
            error: function (data, textStatus, errorThrown) {
                alert(data.responseJSON.Message + "\n" + data.responseJSON.error)
            }
        });
    });
});

