#!/bin/bash

ADDRESS="localhost:9200"
INDEX_NAME="jaeger"

echo
echo "Deleting index '$INDEX_NAME'"
curl -XDELETE "$ADDRESS/$INDEX_NAME?pretty"

echo
echo "Creating index '$INDEX_NAME'"
curl -XPUT "$ADDRESS/$INDEX_NAME?pretty" -H 'Content-Type: application/json' -d'
{
  "settings": {
    "index": {
      "number_of_shards": 1,
      "number_of_replicas": 1
    }
  },
  "mappings": {
    "service": {
      "properties": {
        "service_name": {"type": "text"}
      }
    },
    "operation": {
      "properties": {
        "service_name": {"type": "text"},
        "operation_name": {"type": "text"}
      }
    },
    "span": {
      "properties": {
        "traceID": {"type": "text"},
        "spanID": {"type": "text"},
        "parentSpanID": {"type": "text"},
        "operationName": {"type": "text"},
        "references": {
          "type": "nested",
          "properties": {
            "refType": {"type": "text"},
            "traceID": {"type": "text"},
            "spanID": {"type": "text"}
          }
        },
        "flags": {"type": "integer"},
        "startTime": {"type": "date"},
        "duration": {"type": "long"},
        "tags": {
          "type": "nested",
          "properties": {
            "key": {"type": "text"},
            "vType": {"type": "text"},
            "vStr": {"type": "text"},
            "vNum": {"type": "long"},
            "vBlob": {"type": "binary"}
          }
        },
        "logs": {
          "type": "nested",
          "properties": {
            "message": {"type": "text"},
            "timestamp": {"type": "date"},
            "fields": {
              "type": "nested",
              "properties": {
                "key": {"type": "text"},
                "vType": {"type": "text"},
                "vStr": {"type": "text"},
                "vNum": {"type": "long"},
                "vBlob": {"type": "binary"}
              }
            }
          }
        },
        "process": {
          "type": "object",
          "properties": {
            "serviceName": {"type": "text"},
            "tags": {
              "type": "nested",
              "properties": {
                "key": {"type": "text"},
                "vType": {"type": "text"},
                "vStr": {"type": "text"},
                "vNum": {"type": "long"},
                "vBlob": {"type": "binary"}
              }
            }
          }
        }
      }
    }
  }
}
'
# curl -XPUT "$ADDRESS/$INDEX_NAME?pretty" -H 'Content-Type: application/json' -d'
# {
#   "settings": {
#     "index": {
#       "number_of_shards": 1,
#       "number_of_replicas": 1
#     }
#   },
#   "mappings": {
#     "service": {
#       "properties": {
#         "service_name": {"type": "text"}
#       }
#     },
#     "operation": {
#       "properties": {
#         "service_name": {"type": "text"},
#         "operation_name": {"type": "text"}
#       }
#     },
#     "span": {
#       "properties": {
#         "trace_id": {"type": "long"},
#         "span_id": {"type": "long"},
#         "span_hash": {"type": "long"},
#         "parent_id": {"type": "long"},
#         "operation_name": {"type": "text"},
#         "flags": {"type": "short"},
#         "start_time": {"type": "date"},
#         "duration": {"type": "long"},
#         "tags": {
#           "type": "nested",
#           "properties": {
#             "name": {"type": "text"},
#             "value": {"type": "text"}
#           }
#         },
#         "logs": {
#           "type": "nested",
#           "properties": {
#             "message": {"type": "text"},
#             "fields": {
#               "type": "nested",
#               "properties": {
#                 "name": {"type": "text"},
#                 "value": {"type": "text"}
#               }
#             }
#           }
#         },
#         "refs": {
#           "type": "nested",
#           "properties": {
#             "ref_type": {"type": "text"},
#             "span_id": {"type": "long"}
#           }
#         }
#       }
#     }
#   }
# }
# '

echo "Done"
