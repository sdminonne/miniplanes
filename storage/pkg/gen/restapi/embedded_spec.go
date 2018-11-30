// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "needs to add a description",
    "title": "miniapp storage",
    "version": "1.0.0"
  },
  "paths": {
    "/airlines": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "airlines"
        ],
        "responses": {
          "200": {
            "description": "list of airlines",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/airline"
              }
            }
          },
          "400": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/airports": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "airports"
        ],
        "responses": {
          "200": {
            "description": "list of airports",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/airport"
              }
            }
          },
          "400": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/live": {
      "get": {
        "tags": [
          "liveness"
        ],
        "responses": {
          "200": {
            "description": "liveness probe"
          },
          "503": {
            "description": "if not alive",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/ready": {
      "get": {
        "tags": [
          "readiness"
        ],
        "responses": {
          "200": {
            "description": "readiness probe"
          },
          "503": {
            "description": "if not ready",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/schedules": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "schedules"
        ],
        "responses": {
          "200": {
            "description": "list of schedules",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/schedule"
              }
            }
          },
          "400": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "airline": {
      "type": "object",
      "properties": {
        "Active": {
          "type": "string"
        },
        "AirlineID": {
          "type": "string"
        },
        "Alias": {
          "type": "string"
        },
        "Callsign": {
          "type": "string"
        },
        "Country": {
          "type": "string"
        },
        "IATA": {
          "type": "string"
        },
        "ICAO": {
          "type": "string"
        },
        "Name": {
          "type": "string"
        }
      }
    },
    "airport": {
      "type": "object",
      "properties": {
        "AirportID": {
          "type": "integer"
        },
        "City": {
          "type": "string"
        },
        "Country": {
          "type": "string"
        },
        "DST": {
          "type": "string"
        },
        "IATA": {
          "type": "string"
        },
        "ICAO": {
          "type": "string"
        },
        "Latitude": {
          "type": "number"
        },
        "Longitude": {
          "type": "number"
        },
        "Name": {
          "type": "string"
        },
        "TZ": {
          "type": "string"
        },
        "timezone": {
          "type": "string"
        }
      }
    },
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "schedule": {
      "type": "object",
      "properties": {
        "Arrival": {
          "type": "string"
        },
        "DaysOperated": {
          "type": "string"
        },
        "Departure": {
          "type": "string"
        },
        "Destination": {
          "type": "string"
        },
        "FlightNumber": {
          "type": "string"
        },
        "OperatingCarrier": {
          "type": "string"
        },
        "Origin": {
          "type": "string"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "needs to add a description",
    "title": "miniapp storage",
    "version": "1.0.0"
  },
  "paths": {
    "/airlines": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "airlines"
        ],
        "responses": {
          "200": {
            "description": "list of airlines",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/airline"
              }
            }
          },
          "400": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/airports": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "airports"
        ],
        "responses": {
          "200": {
            "description": "list of airports",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/airport"
              }
            }
          },
          "400": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/live": {
      "get": {
        "tags": [
          "liveness"
        ],
        "responses": {
          "200": {
            "description": "liveness probe"
          },
          "503": {
            "description": "if not alive",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/ready": {
      "get": {
        "tags": [
          "readiness"
        ],
        "responses": {
          "200": {
            "description": "readiness probe"
          },
          "503": {
            "description": "if not ready",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/schedules": {
      "get": {
        "produces": [
          "application/json"
        ],
        "tags": [
          "schedules"
        ],
        "responses": {
          "200": {
            "description": "list of schedules",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/schedule"
              }
            }
          },
          "400": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          },
          "default": {
            "description": "generic error response",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "airline": {
      "type": "object",
      "properties": {
        "Active": {
          "type": "string"
        },
        "AirlineID": {
          "type": "string"
        },
        "Alias": {
          "type": "string"
        },
        "Callsign": {
          "type": "string"
        },
        "Country": {
          "type": "string"
        },
        "IATA": {
          "type": "string"
        },
        "ICAO": {
          "type": "string"
        },
        "Name": {
          "type": "string"
        }
      }
    },
    "airport": {
      "type": "object",
      "properties": {
        "AirportID": {
          "type": "integer"
        },
        "City": {
          "type": "string"
        },
        "Country": {
          "type": "string"
        },
        "DST": {
          "type": "string"
        },
        "IATA": {
          "type": "string"
        },
        "ICAO": {
          "type": "string"
        },
        "Latitude": {
          "type": "number"
        },
        "Longitude": {
          "type": "number"
        },
        "Name": {
          "type": "string"
        },
        "TZ": {
          "type": "string"
        },
        "timezone": {
          "type": "string"
        }
      }
    },
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "schedule": {
      "type": "object",
      "properties": {
        "Arrival": {
          "type": "string"
        },
        "DaysOperated": {
          "type": "string"
        },
        "Departure": {
          "type": "string"
        },
        "Destination": {
          "type": "string"
        },
        "FlightNumber": {
          "type": "string"
        },
        "OperatingCarrier": {
          "type": "string"
        },
        "Origin": {
          "type": "string"
        }
      }
    }
  }
}`))
}
