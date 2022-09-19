import { ChakraProvider, Button, Box, Center, Input, Text, 
    Grid, GridItem, Image, CircularProgress, Tooltip, Link,
    useDisclosure, Drawer, DrawerOverlay,DrawerContent,DrawerHeader,
    DrawerBody } from '@chakra-ui/react';
  import { useEffect, useState } from 'react';
  import { ColorModeSwitcher } from './ColorswitchingComponent';
  
  function Home() {
    const [promptInput, setpromptInput] = useState('');
    const handlePromptInputChange = (event) => setpromptInput(event.target.value)
    const [iterationNums, setiterationNums] = useState("50");
    const handleIterationChange = (event) => setiterationNums(event.target.value)
    const [saveName, setsaveName] = useState("NewFileName_001");
    const handleNameChange = (event) => setsaveName(event.target.value)
  
    var jsonData = {
      PromtString: promptInput,
      PromtIterations: iterationNums,
      PromtSaveName: saveName
    };
  
    
    const [serverIP, setServerIP] = useState(window.location.hostname);
    const handleserverIPInputChange = (event) => setServerIP(event.target.value)

    const { isOpen, onOpen, onClose } = useDisclosure()
  
    var thefilename = "http://"+serverIP+":9998/files/"+saveName+".jpg"
    var postIP = "http://"+serverIP+":9999/test"
  
    
  


    function handleClick() {
      // Send data to the backend via POST
      fetch(postIP, {
        // Enter your IP address here
  
        method: "POST",
        mode: "cors",
        body: JSON.stringify(jsonData) // body data type must match "Content-Type" header
      });
    }
  
    useEffect(() =>{
      fetch(thefilename)}
      ,[]);
  
    return (
      <ChakraProvider>
      <Box  w="100vw" h="100vh">
        <Center>
        <Box w="80%" h="150px">
        <Center>
        <Text marginTop="50px" fontWeight="bold" fontSize="6xl">AIDA</Text>
        </Center>
        </Box></Center>
        <Center>
        <Box margin="0" borderRadius={"lg"} backgroundImage={thefilename} backgroundPosition="center" backgroundRepeat="no-repeat" w="512px" h="512px" maxW={"100%"} ></Box>
        </Center>
        <Center>
        <Box w="80%" h="75px">
        <Center>
        <Input value={promptInput} maxW="50%" width="480px" marginTop="20px" onChange={handlePromptInputChange} placeholder="Enter text prompt here."></Input>
        <Tooltip hasArrow label="Please make sure previous inferencing job is completed before you click generate,else it might cause the program to crash">
        <Button marginLeft="10px" marginTop="20px" onClick={handleClick}>Generate</Button></Tooltip>
        </Center>
        </Box>
        
        </Center>
        
        <Center>
          <Box w="80%" h="75px">
          <Center>
          <Box w="70px">
            <Text w="70px" textAlign={"right"} fontWeight="semibold" fontSize="xs">Iteration</Text>
          </Box>
          <Box  w="45px" ml="5px">
            <Tooltip label="Integer Only. Recommended value is 50, the bigger the number, the longer it took to finish inferencing">
              <Input w="45px" value={iterationNums} onChange={handleIterationChange}  fontWeight="semibold" size="sm"></Input>
            </Tooltip>
          </Box>
          <Box w="60px" ml="10px">
            <Text w="60px"fontSize="xs" fontWeight="semibold">File Name</Text>
          </Box>
          <Box w="100px" ml="5px">
            <Tooltip label="Files will be saved to /Diffusers/examples/inference/ . Please make sure the file name is not the same as existing files, else the program will overwrite the existing file.">
              <Input w="100px" value={saveName} onChange={handleNameChange}fontWeight="semibold" size="sm"></Input>
            </Tooltip>
          </Box>
          
          </Center>
          </Box>
        </Center>
      <Link href="https://ourbunka.com">
        <Box  zIndex="999" position="fixed" top="3%" left="5%">
          <Image borderRadius={"lg"} w="45px" src="https://ourbunka.com:3666/api/ourbunka_logo_4k-01.png"></Image>
        </Box>
      </Link>
      <Button backgroundColor="#09092a" onClick={onOpen} zIndex="999" position="fixed" left="3%" bottom="3%" size={"sm"}><Text color={"pink.100"} fontSize={"xs"} fontWeight="bold">ToS & LICENSE</Text></Button>
      <Box zIndex="999" position="fixed" right="1%" bottom="3%" w="200px">
      <Text fontSize={"10"} fontWeight="thin"><Tooltip hasArrow label="If you're having issue with loading generated images or GENERATE button is not working, replace the default value to your pc local IP address.">AIDAserver IP</Tooltip></Text><Input size="sm" w="80%" value={serverIP} onChange={handleserverIPInputChange} placeholder="localhost"></Input>
      <ColorModeSwitcher/>
      </Box>
      </Box>
      <Drawer h="80%" placement='bottom' onClose={onClose} isOpen={isOpen}>
        <DrawerOverlay/>
        <DrawerContent backgroundColor="#09092a">
          <DrawerHeader color="gray.100" borderBottomWidth='1px'fontSize={"md"} fontWeight="bold">TERMS OF SERVICES & LICENSE 
          <Button backgroundColor="gray.100" onClick={onClose} ml="10px" size="sm"><Text color="#09092a" fontSize={"xs"} fontWeight="bold">AGREE</Text></Button></DrawerHeader>
          
          <DrawerBody>
          <Text color="gray.100" fontSize={"sm"} fontWeight="bold">BY USING AIDA(AIDADIFFUSION) APP, YOU ARE AGREEING WITH OURBUNKA MOTION AIDA(AIDADIFFUSION) <Tooltip hasArrow label="">
            TERMS OF SERVICES</Tooltip>, LICENSING AGREEMENTS AND 
          <Text color="gray.100" color={"blue.200"}><Link isExternal href="https://huggingface.co/spaces/CompVis/stable-diffusion-license">STABLE DIFFUSION LICENSE.</Link></Text>POWERED BY STABLE DIFFUSION V1-4, UI DESIGNED BY 
          <Link isExternal href="https://ourbunka.com/"><Text color={"blue.200"}>OURBUNKA MOTION</Text></Link></Text>
          <Text color="gray.100" fontWeight={"light"} fontSize={"8"}>You agree not to use AIDA(AIDAdiffusion) app, Stable Diffusion Model or Derivatives of the App or Model:</Text>
          <Text color="gray.100" fontWeight={"light"} fontSize={"8"}>- In any way that violates any applicable national, federal, state, local or international law or regulation;</Text>
          <Text color="gray.100" fontWeight={"light"} fontSize={"8"}>- For the purpose of exploiting, harming or attempting to exploit or harm minors in any way;</Text>
          <Text color="gray.100" fontWeight={"light"} fontSize={"8"}>- To generate or disseminate verifiably false information and/or content with the purpose of harming others;</Text>
          <Text color="gray.100" fontWeight={"light"} fontSize={"8"}>- To generate or disseminate personal identifiable information that can be used to harm an individual;</Text>
          <Text color="gray.100" fontWeight={"light"} fontSize={"8"}>- To defame, disparage or otherwise harass others;</Text>
          <Text color="gray.100" fontWeight={"light"} fontSize={"8"}>- For fully automated decision making that adversely impacts an individual’s legal rights or otherwise creates or modifies a binding, enforceable obligation;</Text>
          <Text color="gray.100" fontWeight={"light"} fontSize={"8"}>- For any use intended to or which has the effect of discriminating against or harming individuals or groups based on online or offline social behavior or known or predicted personal or personality characteristics;</Text>
          <Text color="gray.100" fontWeight={"light"} fontSize={"8"}>- To exploit any of the vulnerabilities of a specific group of persons based on their age, social, physical or mental characteristics, in order to materially distort the behavior of a person pertaining to that group in a manner that causes or is likely to cause that person or another person physical or psychological harm;</Text>
          <Text color="gray.100" fontWeight={"light"} fontSize={"8"}>- For any use intended to or which has the effect of discriminating against individuals or groups based on legally protected characteristics or categories;</Text>
          <Text color="gray.100" fontWeight={"light"} fontSize={"8"}>- To provide medical advice and medical results interpretation;</Text>
          <Text color="gray.100" fontWeight={"light"} fontSize={"8"}>- To generate or disseminate information for the purpose to be used for administration of justice, law enforcement, immigration or asylum processes, such as predicting an individual will commit fraud/crime commitment (e.g. by text profiling, drawing causal relationships between assertions made in documents, indiscriminate and arbitrarily-targeted use).</Text>
          </DrawerBody>
        </DrawerContent>
      </Drawer>
      </ChakraProvider>
      
    );
  }
  
  export default Home;