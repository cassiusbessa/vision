package com.github.cassiusbessa.vision.domain.core.valueobjects;

import com.github.cassiusbessa.vision.domain.core.exceptions.DomainException;
import lombok.extern.slf4j.Slf4j;

@Slf4j
public class EncryptedPassword {
    private final String value;

    public EncryptedPassword(String value) {
        this.value = value;
    }

    public String getValue() {
        return value;
    }

    @Override
    public boolean equals(Object obj) {
        if (obj == null) {
            return false;
        }
        if (getClass() != obj.getClass()) {
            return false;
        }
        final EncryptedPassword other = (EncryptedPassword) obj;
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
