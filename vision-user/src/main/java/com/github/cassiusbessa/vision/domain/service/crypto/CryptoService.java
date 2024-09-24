package com.github.cassiusbessa.vision.domain.service.crypto;

import org.springframework.security.crypto.bcrypt.BCryptPasswordEncoder;
import org.springframework.stereotype.Component;

@Component
public class CryptoService {

    private final BCryptPasswordEncoder bCryptPasswordEncoder;

    public CryptoService() {
        this.bCryptPasswordEncoder = new BCryptPasswordEncoder();
    }

    public String encrypt(String password) {
        return bCryptPasswordEncoder.encode(password);
    }

    public boolean matches(String password, String encryptedPassword) {
        return bCryptPasswordEncoder.matches(password, encryptedPassword);
    }
}
