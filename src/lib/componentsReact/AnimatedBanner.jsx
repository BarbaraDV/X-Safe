import React from 'react'
import { AnimationProvider } from '$lib/componentsReact/hooks/useAnimation'
import MainBanner from './Banner.jsx'

const AnimatedBanner = () => {
    return (
        <AnimationProvider>
            <MainBanner />
        </AnimationProvider>
    )
}

export default AnimatedBanner;