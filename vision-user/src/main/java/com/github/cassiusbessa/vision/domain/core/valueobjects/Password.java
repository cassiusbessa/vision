package com.github.cassiusbessa.vision.domain.core.valueobjects;

import com.github.cassiusbessa.vision.domain.core.exceptions.DomainException;

import lombok.extern.slf4j.Slf4j;

@Slf4j
public class Password {
    private final String value;

    public Password(String value) {
        if (!isValid(value)) {
            log.error("Invalid password {} in common domain", value);
            throw new DomainException("Invalid password");
        }
        this.value = value;
    }

    public String getValue() {
        return value;
    }

    private boolean isValid(String password) {
        return password != null && password.length() >= 8 && password.length() <= 20;
    }

    @Override
    public boolean equals(Object obj) {
        if (obj == null) {
            return false;
        }
        if (getClass() != obj.getClass()) {
            return false;
        }
        final Password other = (Password) obj;
        return this.value.equals(other.value);
    }

    @Override
    public int hashCode() {
        return value.hashCode();
    }

    @Override
    public String toString() {
        return value;
    }
}
