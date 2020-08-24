#pragma once

#include <stdint.h>
#include <stdlib.h>
#include <stdio.h>

typedef struct {
    const char* content;
    uint32_t len;
} tc_string_t;

typedef struct  {
    tc_string_t result_json;
    tc_string_t error_json;
} tc_response_t;

typedef struct  {
} tc_response_handle_t;

typedef void (*OnResult)(int request_id, tc_string_t result_json, tc_string_t error_json, int flags);

#ifdef __cplusplus
extern "C" {
#endif

uint32_t tc_create_context();
void tc_destroy_context(uint32_t context);
void tc_json_request_async(uint32_t context, tc_string_t method, tc_string_t params_json, int request_id, OnResult on_result);
void tc_destroy_json_response(const tc_response_handle_t* handle);
tc_response_handle_t* tc_json_request(uint32_t context, tc_string_t method, tc_string_t params_json);
tc_response_t tc_read_json_response(const tc_response_handle_t* handle);

#ifdef __cplusplus
}
#endif




