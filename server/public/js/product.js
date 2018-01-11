productID = 1;

$(document).ready(function () {
    productID = parseInt($("#product_ID").val());


    $("#product_next").click(function () {
        productID++;
        window.location.href = productID;
    });

    $("#product_previous").click(function () {
        productID--;
        window.location.href = productID;
    });

    $("#product_ID").bind('change', function () {
        prodid = parseInt($("#product_ID").val());

        if (prodid > 0) {
            productID++;
            window.location.href = prodid;
        }
    });

    $("#product_delete").click( function(e) {
        e.preventDefault();
        console.log("delete");

        $.ajax({
            url: "/v1/product/" + productID,
            dataType: 'json',
            type: 'DELETE',
            success: function (data) {
                alert(data.Message)
                location.reload()
            },
            error: function (data, textStatus, errorThrown) {
                alert(data.responseJSON.Message + "\n" + data.responseJSON.error)
            }
        })
    })


    $("#product_updateform").submit( function(e) {
        e.preventDefault();

        var data = {};
        data.product_id = productID
        data.product_isbn = $("#product_isbn").val()
        data.product_publisher = $("#product_publisher").val()
        data.product_author = $("#product_author").val()
        data.product_genre = $("#product_genre").val()
        data.product_description = $("#product_description").val()
        data.product_name = $("#product_name").val()
        data.product_type = $("#product_type").val()


        $.ajax({
            url: "/v1/product/" + productID,
            dataType: 'json',
            type: 'PATCH',
            contentType: 'application/json',
            data: JSON.stringify(data),
            success: function (data) {
                location.reload()
            },
            error: function (data, textStatus, errorThrown) {
                alert(data.responseJSON.Message + "\n" + data.responseJSON.error)
            }
        });
    });

    
});

function changeProduct(id) {
    productID = id;
    $("#product_ID").val(id)

    if ( id > 0) {

        $.get("/v1/product/" + productID, function (data) {
            console.log("still",data);
            if (data !== null) {
                data = data.data
                $("#product_ISBN").val(data.isbn)
                $("#product_author").val(data.product_author)
                $("#product_genre").val(data.product_genre)
                $("#product_description").val(data.product_description)
                $("#product_name").val(data.product_name)
                $("#product_type").val(data.product_type)
                $("#product_publisher").val(data.product_publisher)
            }
        })
    }
}
