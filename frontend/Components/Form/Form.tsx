'use client'

import { FormControl, FormLabel, Input, InputGroup,InputLeftAddon,  InputRightElement } from '@chakra-ui/react'

export const Form = () => {
    return (
        <>
            <FormControl>
                <FormLabel>Number of Employees</FormLabel>
                <Input />
                <FormLabel>Annual CO2 Emissions</FormLabel>
                <Input />
                <FormLabel>Annual Electricty Consumption (kwH)</FormLabel>
                <Input />
                <FormLabel>Annual Water Consumption (kL)</FormLabel>
                <Input />
            </FormControl>

        </>
    );
};