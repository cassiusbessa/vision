package com.github.cassiusbessa.vision.domain.service.ports.input;

import java.util.UUID;

public interface TokenService {
    UUID getAccountId(String token);
}
