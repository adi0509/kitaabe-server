USER ROUTE
	- insert new user
	    POST "/api/user"
            {
                "name": "1",
                "email": "1",
                "password": "1",
                "university": "1",
            }
	- get the particular user data
	    GET "/api/user/:id"
	- login
	    POST "/api/user/login"
            {
                "email": "1",
                "password": "1",
            }
	- update user (by email)
	    POST "/api/user/update"
            {
                "name": "1",
                "email": "1",
                "password": "1",
                "university": "1",
            }

ITEM ROUTE
	- insert new item
	    POST "/api/item"
            {
                "item_name": "1",
                "item_description": "1",
                "price": "1",
                "seller_id": "1",
                "available_in_city": "1",
                "category_id": "1",
                "subcategory_id": "1",
                "status": "1",
                "university": "1",
            }
	- get the particular item with id
	    GET "/api/item/id/:id"
	- get the particular item with name
	    GET "/api/item/name/:name"
	- update item
	    POST "/api/item/update"
            {
                "_id": "61922e94d70a055596c93677"
                "item_name": "1",
                "item_description": "1",
                "price": "1",
                "seller_id": "1",
                "available_in_city": "1",
                "category_id": "1",
                "subcategory_id": "1",
                "status": "1",
                "university": "1",
            }

CATEGORY ROUTE
	- insert new category
	    POST "/api/category"
            {
                "category_name": "hello update"
            }
	- get the particular category with id
	    GET "/api/category/id/:id"
	- get the particular category with name
	    GET "/api/category/name/:name"
	- update item
	    POST "/api/category/update"
            {
                "_id": "619516e22a5cec25d461fc6a",
                "category_name": "hello update"
            }

SUBCATEGORY ROUTE
	- insert new item
	    POST "/api/subcategory"
            {
                "category_id": "619516e22a5cec25d461fc6a",
                "subcategory_name": "hello update"
            }
	- get the particular subcategory with id
	    GET "/api/subcategory/id/:id"
	- get the particular subcategory with name
	    GET "/api/subcategory/name/:name"
	- update item
	    POST "/api/subcategory/update"
            {
                "_id": "123"
                "category_id": "619516e22a5cec25d461fc6a",
                "subcategory_name": "hello update"
            }

Media ROUTE
	- insert new item
	    POST "/api/media"
            {
                "item_id": "61922e94d70a055596c93677",
                "url": "googl.com"
            }
	- get the particular media with id
	    GET "/api/media/mediaid/:id"
	- get the particular media with name
	    GET "/api/media/itemid/:id"
	- update item
	    POST "/api/media/update"
            {
                "_id": "123"
                "item_id": "61922e94d70a055596c93677",
                "url": "googl.com"
            }
