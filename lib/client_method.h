#pragma once

#include <stdint.h>
#include <stdbool.h>
#include <stdlib.h>
#include <stdio.h>

// typedef struct {
//     const char* content;
//     uint32_t len;
// } tc_string_t;

// typedef struct  {
//     tc_string_t result_json;
//     tc_string_t error_json;
// } tc_response_t;

// typedef struct  {
// } tc_response_handle_t;

// typedef void (*OnResult)(int request_id, tc_string_t result_json, tc_string_t error_json, int flags);

// #ifdef __cplusplus
// extern "C" {
// #endif

// uint32_t tc_create_context();
// void tc_destroy_context(uint32_t context);
// void tc_json_request_async(uint32_t context, tc_string_t method, tc_string_t params_json, int request_id, OnResult on_result);
// void tc_destroy_json_response(const tc_response_handle_t* handle);
// tc_response_handle_t* tc_json_request(uint32_t context, tc_string_t method, tc_string_t params_json);
// tc_response_t tc_read_json_response(const tc_response_handle_t* handle);

// #ifdef __cplusplus
// }
// #endif

typedef struct {
    const char* content;
    uint32_t len;
} tc_string_data_t;

typedef struct  {
} tc_string_handle_t;

enum tc_response_types_t {
    tc_end = 0,
    tc_response_success = 1,
    tc_response_error = 2,
    tc_response_custom = 100,
};

typedef void (*tc_response_handler_t)(
    uint32_t request_id,
    tc_string_data_t params_json,
    uint32_t response_type,
    bool finished);

#ifdef __cplusplus
extern "C" {
#endif

tc_string_handle_t* tc_create_context(tc_string_data_t config);
void tc_destroy_context(uint32_t context);
void tc_request(
    uint32_t context,
    tc_string_data_t function_name,
    tc_string_data_t function_params_json,
    uint32_t request_id,
    tc_response_handler_t response_handler);

tc_string_handle_t* tc_request_sync(
    uint32_t context,
    tc_string_data_t function_name,
    tc_string_data_t function_params_json);

tc_string_data_t tc_read_string(const tc_string_handle_t* handle);
void tc_destroy_string(const tc_string_handle_t* handle);

#ifdef __cplusplus
}
#endif



