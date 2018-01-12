productID = 1;

$(document).ready(function () {
    productID = parseInt($("#inventory_ID").val());

    $("#product_next").click(function () {
        productID++;
        populateProduct(productID)
    });

    $("#product_previous").click(function () {
        if (productID > 1) {
            productID--;
            populateProduct(productID)
        }
    });

    $("#product_ID").bind('change', function () {
        prodid = parseInt($("#product_ID").val());


        if (prodid > 0) {
            populateProduct(prodid)
        }
    });


    $("#inventory_updateform").submit( function(e) {
        e.preventDefault();
        var data = {};
        data.inventory_id = productID
        data.product_id = productID
        data.inventory_condition = parseInt($("#inventory_condition").val())
        data.inventory_amount = parseInt($("#inventory_amount").val())
        data.inventory_price = parseFloat($("#inventory_price").val())
        data.inventory_note = $("#inventory_note").val()

        $.ajax({
            url: "/v1/inventory",
            dataType: 'json',
            type: 'POST',
            contentType: 'application/json',
            data: JSON.stringify(data),
            success: function (data) {
                window.location.href="/inventory/get/"+data.id
            },
            error: function (data, textStatus, errorThrown) {
                alert(data.responseJSON.Message + "\n" + data.responseJSON.error)
            }
        });
    });

    $("#inventory_condition").val(inventoryCondtion)

    populateProduct(1)
});

function populateProduct( id ) {
    $.get("/v1/product/"+id, function(data) {
        if (data !== null) {
            productID = id
            $("#product_ID").val(id)

            $("#inventory_isbn").val(data.isbn)
            $("#inventory_author").val(data.product_author)
            $("#inventory_genre").val(data.product_genre)
            $("#inventory_publisher").val(data.product_publisher)
            $("#inventory_description").val(data.product_description)
            $("#inventory_name").val(data.product_name)
            $("#inventory_type").val(data.product_type)
        }
    })
}