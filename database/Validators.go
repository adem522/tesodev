package database

import (
	"go.mongodb.org/mongo-driver/bson"
)

var addressSchema = bson.M{
	"bsonType": "object",
	"required": []string{"city", "cityCode", "country"},
	"properties": bson.M{
		"addressLine": bson.M{
			"bsonType":    "string",
			"description": "must be a string",
		},
		"city": bson.M{
			"bsonType":    "string",
			"description": "must be a string and is required",
		},
		"country": bson.M{
			"bsonType":    "string",
			"description": "must be a string and is required",
		},
		"cityCode": bson.M{
			"bsonType":    "int",
			"description": "must be a string and is required",
		},
	},
}

var addressValidator = bson.M{
	"$jsonSchema": addressSchema,
}

var productSchema = bson.M{
	"bsonType": "object",
	"required": []string{"_id", "imageUrl", "name"},
	"properties": bson.M{
		"_id": bson.M{
			"bsonType":    "binData",
			"description": "must be a binData and is required",
		},
		"imageUrl": bson.M{
			"bsonType":    "string",
			"description": "must be a string and is required",
		},
		"name": bson.M{
			"bsonType":    "string",
			"description": "must be a string and is required",
		},
	},
}

var productValidator = bson.M{
	"$jsonSchema": productSchema,
}

var customerSchema = bson.M{
	"bsonType": "object",
	"required": []string{"_id", "name", "email", "address", "createdAt", "updatedAt"},
	"properties": bson.M{
		"_id": bson.M{
			"bsonType":    "binData",
			"description": "must be a binData and is required",
		},
		"name": bson.M{
			"bsonType":    "string",
			"description": "must be a string and is required",
		},
		"email": bson.M{
			"bsonType":    "string",
			"description": "must be a string and is required",
		},
		"address": bson.M{
			"bsonType": "object",
			"required": []string{"city", "country", "cityCode"},
			"properties": bson.M{
				"addressLine": bson.M{
					"bsonType":    "string",
					"description": "must be a string",
				},
				"city": bson.M{
					"bsonType":    "string",
					"description": "must be a string and is required",
				},
				"country": bson.M{
					"bsonType":    "string",
					"description": "must be a string and is required",
				},
				"cityCode": bson.M{
					"bsonType":    "int",
					"description": "must be a int and is required",
				},
			},
			"description": "must be a object and is required",
		},
		"createdAt": bson.M{
			"bsonType":    "date",
			"description": "must be a date and is required",
		},
		"updatedAt": bson.M{
			"bsonType":    "date",
			"description": "must be a date and is required",
		},
	},
}

var customerValidator = bson.M{
	"$jsonSchema": customerSchema,
}

var orderSchema = bson.M{
	"bsonType": "object",
	"required": []string{"_id", "customerId", "quantity", "price", "status", "address", "product", "createdAt", "updatedAt"},
	"properties": bson.M{
		"_id": bson.M{
			"bsonType":    "binData",
			"description": "must be a binData and is required",
		},
		"customerId": bson.M{
			"bsonType":    "binData",
			"description": "must be a string and is required",
		},
		"quantity": bson.M{
			"bsonType":    "int",
			"description": "must be a string and is required",
		},
		"price": bson.M{
			"bsonType":    "double",
			"description": "must be a object and is required",
		},
		"status": bson.M{
			"bsonType":    "string",
			"description": "must be a string and is required",
		},
		"address": bson.M{
			"bsonType": "object",
			"required": []string{"city", "country", "cityCode"},
			"properties": bson.M{
				"addressLine": bson.M{
					"bsonType":    "string",
					"description": "must be a string",
				},
				"city": bson.M{
					"bsonType":    "string",
					"description": "must be a string and is required",
				},
				"country": bson.M{
					"bsonType":    "string",
					"description": "must be a string and is required",
				},
				"cityCode": bson.M{
					"bsonType":    "int",
					"description": "must be a int and is required",
				},
			},
			"description": "must be a object and is required",
		},
		"product": bson.M{
			"bsonType": "object",
			"required": []string{"_id", "imageUrl", "name"},
			"properties": bson.M{
				"_id": bson.M{
					"bsonType":    "binData",
					"description": "must be a binData and is required",
				},
				"imageUrl": bson.M{
					"bsonType":    "string",
					"description": "must be a string and is required",
				},
				"name": bson.M{
					"bsonType":    "string",
					"description": "must be a string and is required",
				},
			},
			"description": "must be a object and is required",
		},
		"createdAt": bson.M{
			"bsonType":    "date",
			"description": "must be a date and is required",
		},
		"updatedAt": bson.M{
			"bsonType":    "date",
			"description": "must be a date and is required",
		},
	},
}

var orderValidator = bson.M{
	"$jsonSchema": orderSchema,
}
