# Google Sheets

Interact with Google Sheets API

# Contents

- [Installation](#Installation)
- [Definitions](#Definitions)
  
  - [Tasks](#Tasks)
    - [get](#get)

# Installation

## MESG Core

This service requires [MESG Core](https://github.com/mesg-foundation/core) to be installed first.

You can install MESG Core by running the following command or [follow the installation guide](https://docs.mesg.com/guide/start-here/installation.html).

```bash
bash <(curl -fsSL https://mesg.com/install)
```

## Service

Download the source code of this service, and then in the service's folder, run the following command:
```bash
mesg-core service deploy
```

# Definitions


# Tasks

## get

Task key: `get`



### Inputs

| **Name** | **Key** | **Type** | **Description** |
| --- | --- | --- | --- |
| **range** | `range` | `String` | Specify range to access to a spesific sheet page, rows and columns |
| **serviceAccountCredentials** | `serviceAccountCredentials` | `String` | Service account credentials to access to Google Sheets API |
| **spreadsheetID** | `spreadsheetID` | `String` | ID of the spreadsheet |

### Outputs

#### error

Output key: `error`



| **Name** | **Key** | **Type** | **Description** |
| --- | --- | --- | --- |
| **message** | `message` | `String` | Error message |

#### success

Output key: `success`



| **Name** | **Key** | **Type** | **Description** |
| --- | --- | --- | --- |
| **rows** | `rows` | `Object` | Rows returned for the query |


