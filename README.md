# decompose_url

Commandline utility to unescape and breakdown unwieldy URLs.

## Install

```bash
  $ go install github.com/rwtnorton/decompose_url
```

## Usage

```bash
  $ decompose_url 'https://github.com/search?q=lacinia'
  https://github.com                          # <-- base url
    /search                                   # <-- path (escaped)
      "q" => "lacinia"                        # <-- query params
```

Or more than one:
```bash
  $ decompose_url 'https://github.com/search?q=lacinia' \
                  'https://slashdot.org/story/02/10/18/2255203/gnarly-error-messages?sbsrc=thisday' \
                  'https://xkcd.com/224/?foo=bar&foo=quux'
  https://github.com
    /search
      "q" => "lacinia"
  https://slashdot.org
    /story/02/10/18/2255203/gnarly-error-messages
      "sbsrc" => "thisday"
  https://xkcd.com
    /224/
      "foo" => ["bar" "quux"]
```

## Author

Richard W. Norton (rwtnorton@gmail.com)
