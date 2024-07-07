package com.github.cassiusbessa.vision.domain.service.exceptions;

public class ResourceAlredyExistsException extends RuntimeException{

    public ResourceAlredyExistsException(String message) {
        super(message);
    }

    public ResourceAlredyExistsException(String message, Throwable cause) {
        super(message, cause);
    }
}
