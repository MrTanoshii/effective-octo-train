import os
from dotenv import load_dotenv
from zeep import Client

load_dotenv()

SOAP_API_WSDL = os.getenv('SOAP_API_WSDL')

client = Client(SOAP_API_WSDL)
print(client.service.__dict__)