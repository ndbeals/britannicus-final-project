customerID = 1;

$(document).ready(function () {

    customerID = parseInt($("#customer_ID").val());


    $("#customer_next").click(function () {
        customerID++;
        window.location.href = customerID;
    });

    $("#customer_previous").click(function () {
        if (customerID > 1) {
            customerID--;
            window.location.href = customerID;
        }
    });

    $("#customer_ID").bind('change enter', function () {
        prodid = parseInt($("#customer_ID").val());

        if (prodid > 0) {
            customerID++;
            window.location.href = prodid;
        }
    });

    $("#customer_delete").click(function (e) {
        e.preventDefault();

        $.ajax({
            url: "/v1/customer/" + customerID,
            dataType: 'json',
            type: 'DELETE',
            success: function (data) {
                alert(data.Message)
                location.reload()
            },
            error: function (data, textStatus, errorThrown) {
                alert(data.responseJSON.Message + "\n" + data.responseJSON.error)
            }
        });
    })


    $("#customer_updateform").submit(function (e) {
        e.preventDefault();

        var data = {};
        data.customer_id = customerID
        data.customer_firstname = $("#customer_firstname").val()
        data.customer_lastname = $("#customer_lastname").val()
        data.customer_email = $("#customer_email").val()
        data.customer_phonenumber = $("#customer_phonenumber").val()
        data.customer_address = $("#customer_address").val()
        data.customer_city = $("#customer_city").val()
        data.customer_state = $("#customer_state").val()
        data.customer_country = $("#customer_country").val()

        $.ajax({
            url: "/v1/customer/" + customerID,
            dataType: 'json',
            type: 'PATCH',
            contentType: 'application/json',
            data: JSON.stringify(data),
            success: function (data) {
                // console.log("DATA POSTED SUCCESSFULLY" , data);
                location.reload()
            },
            error: function (data, textStatus, errorThrown) {
                alert(data.responseJSON.Message + "\n" + data.responseJSON.error)
            }
        });
    });


});