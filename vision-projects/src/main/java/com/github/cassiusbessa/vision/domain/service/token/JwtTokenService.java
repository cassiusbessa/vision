package com.github.cassiusbessa.vision.domain.service.token;

import com.auth0.jwt.JWT;
import com.auth0.jwt.algorithms.Algorithm;
import com.auth0.jwt.interfaces.DecodedJWT;
import com.github.cassiusbessa.vision.domain.service.exceptions.UnauthorizedException;
import com.github.cassiusbessa.vision.domain.service.ports.input.TokenService;
import org.springframework.stereotype.Component;

import java.util.Date;
import java.util.UUID;

@Component
public class JwtTokenService implements TokenService {

    private static final String SECRET = "your-secure";

    private static final long ONE_WEEK_EXPIRATION_TIME = 604800000L;

    public String generateToken(UUID accountId) {
        Algorithm algorithm = Algorithm.HMAC256(SECRET);
        return JWT.create()
                .withIssuer("auth0")
                .withClaim("accountId", accountId.toString())
                .withExpiresAt(new Date(System.currentTimeMillis() + ONE_WEEK_EXPIRATION_TIME))
                .sign(algorithm);
    }

    private boolean validateToken(String token) {
        Algorithm algorithm = Algorithm.HMAC256(SECRET);
        try {
            JWT.require(algorithm).build().verify(token);
            return true;
        } catch (Exception exception) {
            return false;
        }
    }

    public UUID getAccountId(String token) {

        if (!validateToken(token)) {
            throw new UnauthorizedException("Invalid token");
        }
        DecodedJWT decodedJWT = JWT.decode(token);
        return UUID.fromString(decodedJWT.getClaim("accountId").asString());
    }


}
