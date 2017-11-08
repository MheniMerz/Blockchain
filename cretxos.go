package main

/*
This software was developed by employees of the National Institute of 
Standards and Technology (NIST), an agency of the Federal Government. 
Pursuant to title 17 United States Code Section 105, works of NIST 
employees are not subject to copyright protection in the United States 
and are considered to be in the public domain. Permission to freely 
use, copy, modify, and distribute this software and its documentation 
without fee is hereby granted, provided that this notice and disclaimer 
of warranty appears in all copies.

THE SOFTWARE IS PROVIDED 'AS IS' WITHOUT ANY WARRANTY OF ANY KIND, 
EITHER EXPRESSED, IMPLIED, OR STATUTORY, INCLUDING, BUT NOT LIMITED TO, 
ANY WARRANTY THAT THE SOFTWARE WILL CONFORM TO SPECIFICATIONS, ANY 
IMPLIED WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, 
AND FREEDOM FROM INFRINGEMENT, AND ANY WARRANTY THAT THE DOCUMENTATION 
WILL CONFORM TO THE SOFTWARE, OR ANY WARRANTY THAT THE SOFTWARE WILL BE 
ERROR FREE. IN NO EVENT SHALL NIST BE LIABLE FOR ANY DAMAGES, INCLUDING, 
BUT NOT LIMITED TO, DIRECT, INDIRECT, SPECIAL OR CONSEQUENTIAL DAMAGES, 
ARISING OUT OF, RESULTING FROM, OR IN ANY WAY CONNECTED WITH THIS SOFTWARE, 
WHETHER OR NOT BASED UPON WARRANTY, CONTRACT, TORT, OR OTHERWISE, WHETHER 
OR NOT INJURY WAS SUSTAINED BY PERSONS OR PROPERTY OR OTHERWISE, AND 
WHETHER OR NOT LOSS WAS SUSTAINED FROM, OR AROSE OUT OF THE RESULTS OF, 
OR USE OF, THE SOFTWARE OR SERVICES PROVIDED HEREUNDER.
*/

/*
  Stephen Nightingale
  night@nist.gov
  NIST, Information Technology Laboratory
  March 13, 2017
*/


import (
  "flag"
  "fmt"
  "os"
  "currency/methods"
  "currency/structures"
  // "time"
)


// Sum coins created in the ledger.
func main() {

  flag.Parse()
  flay := flag.Args()
  if len(flay) != 1 {
    fmt.Println("Usage: cretxos ledger/newtxes")
    os.Exit(1)
  } // endif flay.

    ledger := []structures.Transaction{}
    ledger = methods.LoadLedger(flay[0], ledger)
    totval := 0
    for ix := 0; ix < len(ledger); ix++ {
      tran := ledger[ix]
      if tran.Ttyp == "CreateCoins" {
        fmt.Printf("%013d/%02d %s %8d %8d %8d %8d %8s\n", tran.Tid, ix, tran.Ttyp, len(tran.Inputs), structures.CoinValues(tran.Inputs), len(tran.Outputs), structures.CoinValues(tran.Outputs), tran.Outputs[0].Owner[:5])
        totval += structures.CoinValues(tran.Outputs)
      } // endif CreateCoins.
    } // end for.
    fmt.Printf("Total Created Value: %d\n", totval)
} // end main.


