#version 410 
layout (location = 0) in vec2 vecPosition;
layout (location = 1) in vec3 vecColour;

out vec3 vertexColour;
void main() {
    gl_Position = vec4(vecPosition, 0.0,1.0);
	vertexColour = vecColour;
}
