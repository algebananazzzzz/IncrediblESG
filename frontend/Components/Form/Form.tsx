'use client'

import { Select, FormControl, FormLabel, Input, InputGroup,InputLeftAddon,  InputRightElement, VStack, Box, Button } from '@chakra-ui/react'
import { Container } from 'postcss'

export const Form = () => {
    return (
        <>
            <FormControl>             
                    <VStack 
                      spacing={8}
                      px={20}
                      paddingRight={1400}
                      align='stretch'
                    >
                        <FormLabel>Sector of Company</FormLabel>
                        <Select placeholder='Select sector...:'>
                            <option value='option1'>Electricity and heat</option>
                            <option value='option2'>Transport</option>s
                            <option value='option3'>Manufacturing and construction</option>
                            <option value='option3'>Buildings</option>
                            <option value='option3'>Industry</option>
                        </Select>
                        <Box>
                            <FormLabel>Number of Employees</FormLabel>
                            <Input />
                        </Box>
                        <Box>
                            <FormLabel>Annual CO2 Emissions (t)</FormLabel>
                            <Input />
                        </Box>
                        <Box>
                            <FormLabel>Annual Electricty Consumption (kwH)</FormLabel>
                            <Input />
                        </Box>
                        <Box>
                            <FormLabel>Annual Water Consumption (kL)</FormLabel>
                            <Input />
                        </Box>
                        <Button colorScheme='blue' >Submit</Button>
                    </VStack>
                        

            </FormControl>

        </>
    );
};