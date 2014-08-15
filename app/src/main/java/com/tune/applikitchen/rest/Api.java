package com.tune.applikitchen.rest;


import com.tune.applikitchen.rest.responses.LoginResponse;

import ly.apps.android.rest.client.Callback;
import ly.apps.android.rest.client.annotations.FormField;
import ly.apps.android.rest.client.annotations.POST;
import ly.apps.android.rest.client.annotations.RestService;

@RestService
public interface Api {

    @POST("/login")
    void authenticate(@FormField("email") String email, @FormField("password") String password, Callback<LoginResponse> callback);


}