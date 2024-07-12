#version 410
out vec4 fragColour;

in vec3 vertexColour;


void main() {
	fragColour = vec4(vertexColour,1);
}