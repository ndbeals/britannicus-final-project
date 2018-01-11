customerID = 1;

$(document).ready(function () {

    customerID = parseInt($("#customer_ID").val());
    console.log(customerID);


    $("#customer_next").click(function () {
        customerID++;
        window.location.href = customerID;
    });

    $("#customer_previous").click(function () {
        if (customerID>1) {
            customerID--;
            window.location.href = customerID;
        }
    });

    $("#customer_ID").bind('change', function () {
        prodid = parseInt($("#customer_ID").val());

        console.log("prodid")

        if (prodid > 0) {
            // ipcRenderer.send("set_inventory_page", pagenum);
            // changecustomer(prodid)
            customerID++;
            window.location.href = prodid;
        }
    });

    $("#customer_delete").click(function (e) {
        e.preventDefault();
        console.log("delete customer");

        $.ajax({
            url: "/v1/customer/" + customerID,
            dataType: 'json',
            type: 'DELETE',
            success: function (data) {
                console.log("DATA POSTED SUCCESSFULLY" , data);
            },
            error: function (jqXhr, textStatus, errorThrown) {
                console.log(errorThrown);
            }
        });
    })


    $("#customer_updateform").submit(function (e) {
        e.preventDefault();
        // $("#customer_ISBN").val(data.isbn)
        // $("#customer_author").val(data.customer_author)
        // $("#customer_genre").val(data.customer_genre)
        // $("#customer_description").val(data.customer_description)
        // $("#customer_name").val(data.customer_name)
        // $("#customer_type").val(data.customer_type)
        console.log("pub", $("#customer_publisher").val());

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


        test = $("#customer_updateform")

        console.log(JSON.stringify(data));


        $.ajax({
            url: "/v1/customer/" + customerID,
            dataType: 'json',
            type: 'PATCH',
            contentType: 'application/json',
            data: JSON.stringify(data),
            success: function (data) {
                console.log("DATA POSTED SUCCESSFULLY" , data);
            },
            error: function (jqXhr, textStatus, errorThrown) {
                console.log(jqXhr, textStatus, errorThrown);
            }
        });
    });


});