package com.tune.applikitchen.rest.responses;

public class LoginResponse {
    private static final String FIELD_TOKEN = "token";

    public String token;
    public boolean status;

    public String getToken() {
        return token;
    }

    public void setToken(String s) {
        token = s;
    }

    public boolean getStatus() {
        return status;
    }

    public void setStatus(boolean b) {
        status = b;
    }
}