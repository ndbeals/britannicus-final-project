productPage = 1;

$(document).ready(function () {

    productPage = parseInt($("#product_page").val());
    console.log(productPage);


    $("#product_next").click(function () {
        if (!changePage(productPage + 1)) {
            changePage(productPage - 1);
        }
    });

    $("#product_previous").click(function () {       
        if (!changePage(productPage - 1)) {
            changePage(productPage+1);
        }
    });

    $("#product_page").bind('keyup input', function () {
        pagenum = parseInt($("#product_page").val());

        if (pagenum > 0) {
            // ipcRenderer.send("set_inventory_page", pagenum);
            oldpage = productPage;
            
            if (!changePage(pagenum)){
                changePage(oldpage);
            }
        }
    });
    // $("#product_page").on("keyup",function () {
    //     console.log("key")
    //     pagenum = $("#product_page").val()

    //     if (pagenum > 0) {
    //         // ipcRenderer.send("set_inventory_page", pagenum);
    //         populateProducts(pagenum);
    //     }
    // });



    $("#productsFilterInput").on("keyup", function () {
        console.log("GA");
        var value = $(this).val().toLowerCase();
        $("#productsFilterTable tr").filter(function () {
            // console.log(this, $(this).index());
            $(this).toggle($(this).text().toLowerCase().indexOf(value) > -1)
        });
    });

    populateProducts(1);




});

function addProduct(table, productID, ISBN, productName, author, genre, publisher, productType, description) {

    var row = "<tr><td>" + 
        productID + "</td><td>" +
        ISBN + "</td><td>" +
        productName + "</td><td>" +
        author + "</td><td>" +
        genre + "</td><td>" +
        publisher + "</td><td>" +
        productType + "</td><td>" +
        description + '</td><td><a href="/product/get/' + productID + '"><button type="button" class="btn btn-primary btn-block tbl-btn">Edit</button></a></td><td><button id="btndel'+productID+'" type="button" class="delbutton btn btn-primary btn-block tbl-btn">Delete</button></td></tr>'
        // </tr>
        // /product/delete/' + productID +
        
    row = table.append(row);

    $('#btndel'+productID).click(function () {
        console.log("Test",productID);
        var myValue = $(this).parent().parent();
        // myValue.remove();
        console.log(myValue);
        console.log(($(this).attr("id")));

        $.get("/product/delete/"+productID, function (data) {
            console.log("succ",data);
            if (data !== null ) {

            }
        }).fail(function (data) {
            //alert("error");
            alert(data.responseJSON.Message)
        });
    });

    return row;
}

function populateProducts( page , hide) {
    $.get("/v1/products/" + page + "/15", function (data) {
        if (data !== null) {
            // productPage = page;
            // $("#product_page").val(page)

            var table = $("#productsFilterTable");

            // table.empty();

            for (var i = 0; i < data.length; i++) {
                item = data[i];

                row = addProduct(table, item.product_id, item.isbn, item.product_name, item.product_author, item.product_genre, item.product_publisher, item.product_type, item.product_description)

                if (hide==true) {
                    changePage(1);
                }
            }

            populateProducts(page+1,true);
        }
        else {
            // populateProducts(productPage);
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
    productPage = page;
    $("#product_page").val(page)
    shown = false;

    $("#productsFilterTable tr").filter(function () {
        index = $(this).index()

        if ((index >= ((page -1) * 15)) && (index < (page *15)) ) {
            $(this).toggle(true);
            shown = true;
        }
        else {
            $(this).toggle(false);
        }
    });

    return shown;
}