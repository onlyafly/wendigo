function pizzaGet(order_id) {
  status = apihub.readKey(order_id);
  apihub.respond(200, '{order_number:@order_id, order_status:@status}');
}

function pizzaPost(request) {
  order_id = apihub.writeNewObject({"name":request.name, "address":request.address, "order":request.order, "phone":request.phone);
  if (checkWithPizzaShop(name,address,order) /*REST API call*/){
    expected_time = notifyPizzaShop(); //REST API call
    apihub.respond(200, {"order_id":order_id, "expected_time": expected_time});
  }
  else {
    apihub.respond(500, {"error":"Pizza shop says no");
  }
}

var endpoint  = { path:"/orders",
                  method:"POST",
                  thing:pizzaPost,
                  request_schema:post-request, // optional
                  response_schema:post-response // optional
                }

var endpoint2 = { path:"/orders/@order_id",
                  method:"GET",
                  thing:pizzaGet,
                  request_schema:get-request, // optional
                  response_schema:get-response // optional
                }

apihub.provide([endpoint, endpoint2]);
