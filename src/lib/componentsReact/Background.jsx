import Particles from "react-particles";
import React, { useCallback } from "react";
import { loadFull } from 'tsparticles';
import particlesConfig from '../media/particles-config'


const BackgroundParticles = () => {
    const particlesInit = useCallback(async engine => {
        await loadFull(engine);
    }, []);

    const particlesLoaded = useCallback(async container => {
    }, []);
    return (
        <Particles id="tsparticles"
            init={particlesInit}
            loaded={particlesLoaded}
            options={particlesConfig}
        />

    )
}

export default BackgroundParticles;