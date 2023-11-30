import React from "react";
import ReactDOM from "react-dom";
import Header from "../componentsReact/Header";
import Headlights from "../componentsReact/Headlights";
import BackgroundParticles from "../componentsReact/Background";
import AnimatedBanner from "../componentsReact/AnimatedBanner";

export function renderHeader(container, user) {
    ReactDOM.render(<Header user={user} />, container);
}

export function renderHeadlights(container) {
    ReactDOM.render(<Headlights />, container);
}

export function renderBackground(container) {
    ReactDOM.render(<BackgroundParticles />, container);
}

export function renderBanner(container) {
    ReactDOM.render(<AnimatedBanner />, container);
}

export function unmountReact(container) {
    ReactDOM.unmountComponentAtNode(container);
}
