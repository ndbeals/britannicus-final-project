productPage = 1;

$(document).ready(function () {
    productPage = parseInt($("#product_page").val());


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
            oldpage = productPage;
            
            if (!changePage(pagenum)){
                changePage(oldpage);
            }
        }
    });


    $("#productsFilterInput").on("keyup", function () {
        var value = $(this).val().toLowerCase();

        if (value == "") {
            changePage(productPage);
        } else {
            results = 0
            $("#productsFilterTable tr").filter(function () {
                if ($(this).is(":visible")) {
                    $(this).toggle(false)
                }
                if (results < 15) {
                    show = $(this).text().toLowerCase().indexOf(value) > -1
                    $(this).toggle(show)
                    if (show) {
                        results++;
                    }
                }
            });
        }
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
        description + '</td><td><a href="/product/get/' + productID + '"><button type="button" class="btn btn-primary btn-block tbl-btn">Edit</button></a></td><td><button id="btndel'+productID+'" type="button" class="btn btn-danger btn-block tbl-btn">Delete</button></td></tr>'
        
    row = $(row).appendTo(table)//table.append(row);

    $('#btndel'+productID).click(function () {
        // console.log("Test",productID);
        var parent = $(this).parent().parent();
        // console.log(myValue);
        // console.log(($(this).attr("id")));

        $.get("/product/delete/"+productID, function (data) {
            // console.log("succ",data);
            if (data !== null ) {
                parent.remove();
                // changePage(productPage);
            }
        }).fail(function (data) {
            alert(data.responseJSON.Message)
        });
    });

    return row;
}

function populateProducts( page , hide) {
    $.get("/v1/products/" + page + "/15", function (data) {
        if (data !== null) {
            var table = $("#productsFilterTable");

            // table.empty();

            for (var i = 0; i < data.length; i++) {
                item = data[i];

                row = addProduct(table, item.product_id, item.isbn, item.product_name, item.product_author, item.product_genre, item.product_publisher, item.product_type, item.product_description)

                if (hide==true) {
                    row.toggle(false)
                    // changePage(productPage);
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