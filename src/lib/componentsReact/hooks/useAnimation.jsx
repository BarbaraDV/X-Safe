import React, { createContext, useState, useContext } from 'react';

const AnimationContext = createContext();


export const AnimationProvider = ({ children }) => {
    const [isTitleAnimated, setTitleAnimated] = useState(false);

    return (
        <AnimationContext.Provider value={{ isTitleAnimated, setTitleAnimated }}>
            {children}
        </AnimationContext.Provider>
    );
};


export const useAnimationContext = () => useContext(AnimationContext);