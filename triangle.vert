#version 410 
layout (location = 0) in vec2 vecPosition;
layout (location = 1) in vec3 vecColour;

out vec4 vertexColour;
void main() {
	float d = distance(vec2(0.5,0.5),vecPosition);
    gl_Position = vec4(vecPosition, 0.0,1.0);
	vertexColour = vec4(vecColour,1);
}
