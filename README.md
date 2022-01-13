# conttest

<img src="log.png" width="400">

![build](https://github.com/eigenhombre/conttest/actions/workflows/build.yml/badge.svg)

Continuous testing helper / file watcher, similar in design to 
[this Python version](https://github.com/eigenhombre/continuous-testing-helper),
rewritten in Go for ease of installation and performance.

# Install

Install Go using your OS'es package manager.  Then,

    go install github.com/eigenhombre/conttest@latest
    
# Usage

When `conttest` is run in a directory, its argument will be run once, then each time a file in that directory is saved.

## Examples

    # Simple command: no quotes needed:
    $ conttest ./build.sh

    # Add quotes to pass arguments to the target command...
    $ conttest "make -n"

    # ... or to run compound commands or pipes:
    $ conttest 'make && make install && myprog'
    $ conttest 'go test && say "OK" || say "FAIL"'

# License

Copyright © 2022, John Jacobsen. MIT License.

# Disclaimer

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
