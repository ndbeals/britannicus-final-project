inventoryID = 1;

$(document).ready(function () {

    inventoryID = parseInt($("#inventory_ID").val());

    $("#inventory_next").click(function () {
        inventoryID++;
        window.location.href = inventoryID;
    });

    $("#inventory_previous").click(function () {
        if (inventoryID > 1){
            inventoryID--;
            window.location.href = inventoryID;
        }
    });

    $("#inventory_ID").bind('change', function () {
        prodid = parseInt($("#inventory_ID").val());

        if (prodid > 0) {
            inventoryID++;
            window.location.href = prodid;
        }
    });

    $("#inventory_delete").click( function(e) {
        e.preventDefault();
        console.log("delete");

        $.ajax({
            url: "/v1/inventory/" + inventoryID,
            dataType: 'json',
            type: 'DELETE',
            success: function (data) {
                // console.log("DATA POSTED SUCCESSFULLY" , data);
                alert(data.Message)
                location.reload()
            },
            error: function (data, textStatus, errorThrown) {
                alert(data.responseJSON.Message + "\n" + data.responseJSON.error)
            }
        });
    })


    $("#inventory_updateform").submit( function(e) {
        e.preventDefault();

        var data = {};
        data.inventory_id = inventoryID
        data.inventory_condition = parseInt($("#inventory_condition").val())
        data.inventory_amount = parseInt($("#inventory_amount").val())
        data.inventory_price = parseFloat($("#inventory_price").val())
        data.inventory_note = $("#inventory_note").val()


        $.ajax({
            url: "/v1/inventory/" + inventoryID,
            dataType: 'json',
            type: 'PATCH',
            contentType: 'application/json',
            data: JSON.stringify(data),
            success: function (data) {
                console.log("DATA POSTED SUCCESSFULLY" , data);
            },
            error: function (data, textStatus, errorThrown) {
                alert(data.responseJSON.Message + "\n" + data.responseJSON.error)
            }
        });
    });

    $("#inventory_condition").val(inventoryCondtion)
});