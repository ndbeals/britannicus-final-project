$(document).ready(function () {
    $("#customer_updateform").submit(function (e) {
        e.preventDefault();

        var data = {};
        data.customer_id = -1
        data.customer_firstname = $("#customer_firstname").val()
        data.customer_lastname = $("#customer_lastname").val()
        data.customer_email = $("#customer_email").val()
        data.customer_phonenumber = $("#customer_phonenumber").val()
        data.customer_address = $("#customer_address").val()
        data.customer_city = $("#customer_city").val()
        data.customer_state = $("#customer_state").val()
        data.customer_country = $("#customer_country").val()


        $.ajax({
            url: "/v1/customer",
            dataType: 'json',
            type: 'POST',
            contentType: 'application/json',
            data: JSON.stringify(data),
            success: function (data) {
                window.location.href ="/customer/get/"+data.id
            },
            error: function (data, textStatus, errorThrown) {
                alert(data.responseJSON.Message + "\n" + data.responseJSON.error)
            }
        });
    });


});