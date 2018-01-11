inventoryID = 1;
customerID = 1;
itemList ={};

$(document).ready(function () {

    inventoryID = parseInt($("#inventory_ID").val());
    console.log(inventoryID);


    $("#item_next").click(function () {
        inventoryID++;
        populateItem(inventoryID)
    });

    $("#item_previous").click(function () {
        if (inventoryID > 1){
            inventoryID--;

            populateItem(inventoryID)
        }
    });

    $("#item_add").click(function () {
        console.log("get order item")
        getOrderItem(inventoryID)
    });

    $("#item_clear").click(function () {
        $("#orderlisttable").empty()
        itemList={};
    });

    $("#item_ID").bind('change', function () {
        itemid = parseInt($("#inventory_ID").val());

        if (itemid > 0) {
            populateItem(itemid)
        }
    });


    $("#customer_next").click(function () {
        customerID++;
        populateCustomer(customerID)
    });

    $("#customer_previous").click(function () {
        if (customerID > 1) {
            customerID--;

            populateCustomer(customerID)
        }
    });

    $("#customer_ID").bind('change', function () {
        customerid = parseInt($("#customer_ID").val());

        if (customerid > 0) {
            populateCustomer(customerid)
        }
    });


    $("#order_createform").submit( function(e) {
        e.preventDefault();
        // $("#inventory_ISBN").val(data.isbn)
        // $("#inventory_author").val(data.inventory_author)
        // $("#inventory_genre").val(data.inventory_genre)
        // $("#inventory_description").val(data.inventory_description)
        // $("#inventory_name").val(data.inventory_name)
        // $("#inventory_type").val(data.inventory_type)
        // console.log("pub",$("#inventory_publisher").val());

        var data = {};
        data.customer_id = customerID;
        data.item_list = itemList;
        // data.inventory_condition = parseInt($("#inventory_condition").val())
        // data.inventory_amount = parseInt($("#inventory_amount").val())
        // data.inventory_price = parseFloat($("#inventory_price").val())
        // data.inventory_note = $("#inventory_note").val()
        // data.inventory_description = $("#inventory_description").val()
        // data.inventory_name = $("#inventory_name").val()
        // data.inventory_type = $("#inventory_type").val()
        
        // console.log($("#inventory_condition").val())

        test = $("#inventory_updateform")

        console.log(JSON.stringify(data));


        $.ajax({
            url: "/v1/order",
            dataType: 'json',
            type: 'POST',
            contentType: 'application/json',
            data: JSON.stringify(data),
            success: function (data) {
                console.log("DATA POSTED SUCCESSFULLY" , data);
                // window.location.href="/order/get/"+data.data.order_id
            },
            error: function (jqXhr, textStatus, errorThrown) {
                console.log(errorThrown);
            }
        });
    });

    // $("#inventory_condition").val(inventoryCondtion)

    populateItem(1)
    populateCustomer(1)
});

function populateItem( id ) {
    $.get("/v1/inventory/"+id, function(data) {
        console.log("dataaaa",data)
        if (data !== null) {
            inventoryID = id
            $("#item_ID").val(id)

            $("#item_isbn").val(data.product.isbn)
            $("#item_author").val(data.product.product_author)
            $("#item_genre").val(data.product.product_genre)
            $("#item_publisher").val(data.product.product_publisher)
            $("#item_description").val(data.product.product_description)
            $("#item_name").val(data.product.product_name)
            $("#item_type").val(data.product.product_type)
            $("#inventory_price").val(data.item_price)
        }
    })
}

function populateCustomer(id) {
    $.get("/v1/customer/" + id, function (data) {
        console.log("cutst",data)
        if (data !== null) {
            customerID = id
            $("#customer_ID").val(id)

            $("#customer_firstname").val( data.first_name ) 
            $("#customer_lastname").val( data.last_name ) 
            $("#customer_email").val( data.customer_email ) 
            $("#customer_phonenumber").val( data.customer_phone ) 
            $("#customer_address").val( data.customer_address ) 
            $("#customer_city").val( data.customer_city ) 
            $("#customer_state").val( data.customer_state ) 
            $("#customer_country").val( data.customer_country ) 

            // $("#item_isbn").val(data.product.isbn)
            // $("#item_author").val(data.product.product_author)
            // $("#item_genre").val(data.product.product_genre)
            // $("#item_publisher").val(data.product.product_publisher)
            // $("#item_description").val(data.product.product_description)
            // $("#item_name").val(data.product.product_name)
            // $("#item_type").val(data.product.product_type)
        }
    })
}

function updateList() {
    $.get("/v1/order/" + orderID, function (data) {
        if (data !== null) {
            table = $("#orderlisttable")

            for (var i = 0; i < data.item_list.length; i++) {

                addOrderItem(table, orderID, data.item_list[i])
            }
        }
    })
}

function getOrderItem( id ) {
    $.get("/v1/inventory/" + id, function (data) {
        console.log("orderitema", data)
        if (data !== null) {

            amount = parseInt($("#order_amount").val())
            console.log("amount")
            if (amount > 0 && amount < 100) {
                table = $("#orderlisttable")

                addOrderItem( table ,amount, data )
                
            }
            // inventoryID = id
            // $("#item_ID").val(id)

            // $("#item_isbn").val(data.product.isbn)
            // $("#item_author").val(data.product.product_author)
            // $("#item_genre").val(data.product.product_genre)
            // $("#item_publisher").val(data.product.product_publisher)
            // $("#item_description").val(data.product.product_description)
            // $("#item_name").val(data.product.product_name)
            // $("#item_type").val(data.product.product_type)
        }
    })
}

function addOrderItem(table, amount, item) {
    console.log(item)
    idint = parseInt(item.inventory_id)
    itemList[idint]=amount
    console.log(itemList)
    var row = `
    <tr>
        <td>
            <div>
                `+ item.inventory_id + `
            </div>
        </td>
        <td>
            <div>
               `+ item.product.product_name + `
            </div>
        </td>
        <td>
            <div>
                `+ amount + `
            </div>
        </td>
        <td>
            <div>
                `+ ( parseFloat(item.item_price) * amount) + `
            </div>
        </td>
    </tr>`

    row = $(row).appendTo(table)//table.append(row);
    // row.toggle(false)


    // $('#btndel' + order_id).click(function () {
    //     var parent = $(this).parent().parent();
    //     $.ajax({
    //         url: "/v1/order/" + order_id,
    //         dataType: 'json',
    //         type: 'DELETE',
    //         success: function (data) {
    //             if (data !== null) {
    //                 parent.remove()
    //             }
    //         },
    //         error: function (data, textStatus, errorThrown) {
    //             // alert(data.responseJSON.Message)
    //             console.log(data);
    //         },
    //     });
    // });

    return row;
}