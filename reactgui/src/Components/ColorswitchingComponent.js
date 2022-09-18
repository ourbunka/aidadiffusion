import React from "react";
import {useColorMode, useColorModeValue, IconButton} from '@chakra-ui/react'
import { FaMoon, FaSun} from 'react-icons/fa';

export const ColorModeSwitcher = props => {
    const {toggleColorMode} = useColorMode();
    const text = useColorModeValue('dark', 'light');
    const SwitchIcon = useColorModeValue(FaMoon, FaSun);

    return(
        <IconButton
        _hover={{bg:" rgba(233,212,212,0.2)"}}
        size="sm"
        borderRadius="100"
        fontSize="lg"
        aria-label={`Switch to ${text} mode`}
        variant="ghost"
        color="current"
        onClick={toggleColorMode}
        icon={<SwitchIcon/>}
        {...props}
        />
    )
}