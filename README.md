# FealTY

FealTY - An open source and easy to deploy customer rewards scheme.

Follow the below instructions to get this all set up! (Please report any issues [here](https://github.com/dangeroustech/fealty/issues))

## Outcomes

Once setup is complete, this will create the following infrastructure for you:

- A database server to hold the records of customers reward points (Linode don't have managed DBs yet)
- An app server to accept API requests, interrogate the Database, and provide basic frontend rendering for web browsers
- A Linode Domain with an A record at rewards.domain.xyz pointing at the app server

## Prerequisites

- Linode Account ([Sign Up Here](https://login.linode.com/signup))
- GitHub Account ([Sign Up Here](https://github.com/join))
- A registered domain with Nameservers pointed to Linode (optional) ([Walkthrough Here](https://www.linode.com/docs/guides/dns-manager/))

Once you have accounts to both of these services, this README will walk you through setting up everything else!

## Linode API Key

Follow this [Linode Tutorial](https://www.linode.com/docs/guides/getting-started-with-the-linode-api/#get-an-access-token) to obtain an API Key.

When selecting the scope, 'Select All' for Read/Write will work fine, if you want more granular options, the services that we _need_ are:

- Domains
- Images
- Linodes
- Object Storage
- Volumes

Make sure these have Read/Write permissions and you'll be golden.

## Linode Object Storage

In order for our infrastructure to be deployed from 'the cloud' we need to create a Linode Object Storage bucket to hold the current status of our servers.

### Create The Bucket

- Navigate to the 'Object Storage' tab of your Linode Dashboard and click 'Create Bucket'

![OS1](docs/readme/OS1.png)

- Give your bucket a label (I'd suggest something related to the task such as fealty/rewards) and a region (I'd suggest Atlanta because that's where our servers will be (currently because of Linode VLAN offering) but it doesn't really matter)

![OS2](docs/readme/OS2.png)

- Agree to the warning about a $5 a month charge, you won't exceed any bandwidth caps and you can use these buckets to replace things like Dropbox/Google Drive, if you're looking for extra monetary value

![OS3](docs/readme/OS3.png)

### Create an Access Key

- From within the 'Object Storage' tab, select 'Access Keys' and click 'Create Access Key'

![OS4](docs/readme/OS4.png)

- Give your key a label (again, I'd suggest fealty/rewards), switch on 'Limited Access' and make sure to select the bucket you just created for read/write access (this will look like you're selecting all for now but if you create other buckets this key won't have access, best practise really 🤷‍♂️)

![OS5](docs/readme/OS5.png)

- When the Access Key and Secret pop up, __COPY THESE SOMEWHERE__ (a notepad document will do for now but if you have a password manager that's preferrable)

![OS6](docs/readme/OS6.png)

## GitHub Bits

### Fork The Repo

- When on the main GitHub repository page, press the 'Fork' button in the top-right corner

![GH1](docs/readme/GH1.png)

- Wait for the forking screen to disappear

![GH2](docs/readme/GH2.png)

- You will now notice the repository again, except under your username in the top left corner

![GH3](docs/readme/GH3.png)

### Add Keys

- Click on the 'Settings' tab from the main repository view

![GH4](docs/readme/GH4.png)

- Navigate to 'Secrets' on the left sidebar

![GH5](docs/readme/GH5.png)

- Use the 'New repository secret' button to add the following secrets:

  - LINODE_TOKEN - This is the API token you generated in the first Linode step
  - OBJECT_ACCESS_KEY - This is the Access Key from the Linode Object Storage step
  - OBJECT_SECRET_KEY - This is the Secret Key from the Linode Object Storage step
  - DOMAIN - This is your domain (you don't even have to own it, it just needs to be unique to Linode, but if you don't own it then it won't _actually_ work, ofc)

![GH6](docs/readme/GH6.png)

## Dev Info - API Routes

### Accounts

- [ ] Update (Partial) - **STRETCH**
