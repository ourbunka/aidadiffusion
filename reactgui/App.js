import { ChakraProvider, Button, Box, Center, Input, Text, 
  Grid, GridItem, Image, CircularProgress, Tooltip, Link } from '@chakra-ui/react';
import { useEffect, useState } from 'react';
import { ColorModeSwitcher } from './Components/ColorswitchingComponent';

function App() {
  const [promptInput, setpromptInput] = useState('');
  const handlePromptInputChange = (event) => setpromptInput(event.target.value)
  const [iterationNums, setiterationNums] = useState(50);
  const handleIterationChange = (event) => setiterationNums(event.target.value)
  const [saveName, setsaveName] = useState("NewFileName_001");
  const handleNameChange = (event) => setsaveName(event.target.value)

  var jsonData = {
    PromtString: promptInput,
    PromtIterations: iterationNums,
    PromtSaveName: saveName
  };

  
  const [serverIP, setServerIP] = useState('');
  const handleserverIPInputChange = (event) => setServerIP(event.target.value)

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
    <Box w="100%" h="100%" marginTop={0} marginBottom={0}>
      <Center>
      <Box  w="80%" h="200px">
      <Center>
      <Text marginTop="50px" fontWeight="bold" fontSize="6xl">AIDA</Text>
      </Center>
      </Box></Center>
      <Center>
      <Box  w="80%" h="100px">
      
      <Box  w="100%" h="20px">
      </Box>
      <Center>
      <Input value={promptInput} width="60%" marginTop="20px" onChange={handlePromptInputChange} placeholder="Enter text prompt here."></Input>
      <Tooltip hasArrow label="Please make sure previous inferencing job is completed before you click generate,else it might cause the program to crash">
      <Button marginLeft="10px" marginTop="20px" onClick={handleClick}>Generate</Button></Tooltip>
      </Center>
      </Box>
      </Center>
      <Center>
      <Box w="70%" marginLeft="-25%">
      <Grid  templateColumns='repeat(2, 2fr)' gap={2}>
        <GridItem w='100%' h='10' ><Text textAlign={"right"} fontWeight="semibold">Iteration</Text></GridItem>
        <GridItem w='100%' h='10' >
        <Tooltip label="Integer Only. Recommended value is 50, the bigger the number, the longer it took to finish inferencing">
        <Input value={iterationNums} onChange={handleIterationChange}  fontWeight="semibold" size="sm"></Input>
        </Tooltip></GridItem>
        <GridItem w='100%' h='10' ><Text textAlign={"right"} fontWeight="semibold">File Name</Text></GridItem>
        <GridItem w='100%' h='10' >
        <Tooltip label="Files will be saved to /Diffusers/examples/inference/ . Please make sure the file name is not the same as existing files, else the program will overwrite the existing file.">
        <Input value={saveName} onChange={handleNameChange}fontWeight="semibold" size="sm"></Input>
        </Tooltip></GridItem>
      </Grid>
      </Box>
      </Center>
      <Center>
      <Box backgroundImage={thefilename} w="512px" h="512px" ></Box>
      </Center>
    </Box>
    <Center>
    <Grid width="50%" marginTop="10px" templateColumns='repeat(4, 2fr)' gap={1}>
        <GridItem w='100%' h='10' ><Text marginTop="10px" textAlign={"right"} fontWeight="semibold">Search File</Text></GridItem>
        <GridItem colSpan="3" w='100%' h='10' >
        <Tooltip label="Enter file name(without .jpg extension) to search previous generated images.">
        <Input width="90%" value={saveName} onChange={handleNameChange}></Input>
        </Tooltip></GridItem>
      </Grid>
    </Center>
    <Box zIndex="999" position="fixed" left="1%" bottom="0.5%" maxW="300px"><Text fontSize={"8"} fontWeight="bold">BY USING AIDA(AIDADIFFUSION) APP, YOU ARE AGREEING WITH OURBUNKA MOTION AIDA(AIDADIFFUSION) <Tooltip hasArrow label="You agree not to use AIDA(AIDAdiffusion) app, Stable Diffusion Model or Derivatives of the App or Model:
- In any way that violates any applicable national, federal, state, local or international law or regulation;
- For the purpose of exploiting, harming or attempting to exploit or harm minors in any way;
- To generate or disseminate verifiably false information and/or content with the purpose of harming others;
- To generate or disseminate personal identifiable information that can be used to harm an individual;
- To defame, disparage or otherwise harass others;
- For fully automated decision making that adversely impacts an individual’s legal rights or otherwise creates or modifies a binding, enforceable obligation;
- For any use intended to or which has the effect of discriminating against or harming individuals or groups based on online or offline social behavior or known or predicted personal or personality characteristics;
- To exploit any of the vulnerabilities of a specific group of persons based on their age, social, physical or mental characteristics, in order to materially distort the behavior of a person pertaining to that group in a manner that causes or is likely to cause that person or another person physical or psychological harm;
- For any use intended to or which has the effect of discriminating against individuals or groups based on legally protected characteristics or categories;
- To provide medical advice and medical results interpretation;
- To generate or disseminate information for the purpose to be used for administration of justice, law enforcement, immigration or asylum processes, such as predicting an individual will commit fraud/crime commitment (e.g. by text profiling, drawing causal relationships between assertions made in documents, indiscriminate and arbitrarily-targeted use).">
  TERMS OF SERVICES</Tooltip>, LICENSING AGREEMENTS AND <Text color={"red"}><Link isExternal href="https://huggingface.co/spaces/CompVis/stable-diffusion-license">STABLE DIFFUSION LICENSE.</Link></Text>POWERED BY STABLE DIFFUSION V1-4, UI DESIGNED BY <Link isExternal href="https://ourbunka.com/">OURBUNKA MOTION</Link></Text></Box>

<ColorModeSwitcher zIndex="999" position="fixed" right="3%" bottom="3%"></ColorModeSwitcher>
<Box zIndex="998" position="fixed" w="200px" right="8%" bottom="3%" ><Input value={serverIP} onChange={handleserverIPInputChange} placeholder="localhost"></Input></Box>
</ChakraProvider>
  );
}

export default App;
