package com.tune.applikitchen.rest.responses;

/**
 * Created by johnny on 8/14/14.
 */
public class LoginResponse {
    public String getToken() {
        return Token;
    }

    public void setToken(String token) {
        Token = token;
    }

    private String Token;

    public boolean isStatus() {
        return Status;
    }

    public void setStatus(boolean status) {
        Status = status;
    }

    private boolean Status;
}