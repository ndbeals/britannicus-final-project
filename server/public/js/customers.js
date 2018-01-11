customerPage = 1;

$(document).ready(function () {
    customerPage = parseInt($("#customer_page").val());


    $("#customer_next").click(function () {
        if (!changePage(customerPage + 1)) {
            changePage(customerPage - 1);
        }
    });

    $("#customer_previous").click(function () {       
        if (!changePage(customerPage - 1)) {
            changePage(customerPage+1);
        }
    });

    $("#customer_page").bind('keyup input', function () {
        pagenum = parseInt($("#customer_page").val());

        if (pagenum > 0) {
            oldpage = customerPage;
            
            if (!changePage(pagenum)){
                changePage(oldpage);
            }
        }
    });


    $("#customersFilterInput").bind("keyup inpit", function () {
        var value = $(this).val().toLowerCase();
        
        if (value == "") {
            changePage(customerPage);
        } else {
            results = 0
            $("#customersFilterTable tr").filter(function () {
                if ($(this).is(":visible")) {
                    $(this).toggle(false)
                }
                if ( results < 15 ){
                    show = $(this).text().toLowerCase().indexOf(value) > -1
                    $(this).toggle( show )
                    if (show) {
                        results++;
                    }
                }
            });
        }
    });

    populateCustomers(1);
});

function addCustomer(table, customerID, first_name, last_name, customer_email, customer_phone, customer_address, customer_city, customer_state, customer_country) {

    var row = "<tr><td>" + 
        customerID + "</td><td>" +
        first_name + "</td><td>" +
        last_name + "</td><td>" +
        customer_email + "</td><td>" +
        customer_phone + "</td><td>" +
        customer_address + "</td><td>" +
        customer_city + "</td><td>" +
        customer_state + "</td><td>" +
        customer_country + '</td><td><a href="/customer/get/' + customerID + '"><button type="button" class="btn btn-primary btn-block tbl-btn">Edit</button></a></td><td><button id="btndel'+customerID+'" type="button" class="delbutton btn btn-primary btn-block tbl-btn">Delete</button></td></tr>'
        
    row = $(row).appendTo(table); // table.append(row);
    $('#btndel'+customerID).click(function () {
        // console.log("Test",customerID);
        var parent = $(this).parent().parent();
        // console.log(myValue);
        // console.log(($(this).attr("id")));

        $.get("/customer/delete/"+customerID, function (data) {
            console.log("succ",data);
            if (data !== null ) {
                parent.remove();
                changePage(customerPage);
            }
        }).fail(function (data) {
            alert(data.responseJSON.Message)
        });
    });

    return row;
}

function populateCustomers( page , hide) {
    $.get("/v1/customers/" + page + "/15", function (data) {
        if (data !== null) {
            var table = $("#customersFilterTable");

            // table.empty();

            for (var i = 0; i < data.length; i++) {
                item = data[i];

                row = addCustomer(table, item.customer_id, item.first_name, item.last_name, item.customer_email, item.customer_phone, item.customer_address, item.customer_city, item.customer_state, item.customer_country)

                if (hide==true) {
                    row.toggle(false);
                    // changePage(customerPage);
                }
                if (page==customerPage){
                    changePage(customerPage)
                }
            }

            populateCustomers(page+1,true);
        }
        else {
            // populateCustomers(customerPage);
        }
    })
        .done(function () {
            //alert("second success");
        })
        .fail(function (data) {
            //alert("error");
        })
        .always(function (data) {
            //alert("finished" + data);
        });
}

function changePage(page) {
    customerPage = page;
    $("#customer_page").val(page)
    shown = false;

    $("#customersFilterTable tr").filter(function () {
        index = $(this).index()

        if ((index >= ((page -1) * 15)) && (index < (page * 15)) ) {
            $(this).toggle(true);
            shown = true;
        }
        else {
            $(this).toggle(false);
        }
    });

    return shown;
}