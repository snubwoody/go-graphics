#version 410 

layout (location = 0) in vec2 vecPosition;
layout (location = 1) in vec4 vecColour;

out vec4 vertexColour;

void main() {
	vertexColour = vecColour;
    gl_Position = vec4(vecPosition, 0.0,1.0);
}
