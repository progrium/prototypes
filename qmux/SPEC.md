# qmux

qmux is a wire protocol for multiplexing connections or streams into a single connection.
It's extracted from the SSH Connection Protocol and simplified greatly.

## Channels

   Either side may open a channel.  Multiple channels are multiplexed
   into a single connection.

   Channels are identified by numbers at each end.  The number referring
   to a channel may be different on each side.  Requests to open a
   channel contain the sender's channel number.  Any other channel-
   related messages contain the recipient's channel number for the
   channel.

   Channels are flow-controlled.  No data may be sent to a channel until
   a message is received to indicate that window space is available.

###  Opening a Channel

   When either side wishes to open a new channel, it allocates a local
   number for the channel.  It then sends the following message to the
   other side, and includes the local channel number and initial window
   size in the message.

      byte      QMUX_MSG_CHANNEL_OPEN
      uint64    sender channel
      uint64    initial window size
      uint64    maximum packet size

   The 'sender channel' is a local identifier for the channel used by the
   sender of this message.  The 'initial window size' specifies how many
   bytes of channel data can be sent to the sender of this message
   without adjusting the window. The 'maximum packet size' specifies the
   maximum size of an individual data packet that can be sent to the
   sender.  For example, one might want to use smaller packets for
   interactive connections to get better interactive response on slow
   links.

###  Data Transfer

   The window size specifies how many bytes the other party can send
   before it must wait for the window to be adjusted.  Both parties use
   the following message to adjust the window.

      byte      QMUX_MSG_CHANNEL_WINDOW_ADJUST
      uint64    recipient channel
      uint64    bytes to add

   After receiving this message, the recipient MAY send the given number
   of bytes more than it was previously allowed to send; the window size
   is incremented.  Implementations MUST correctly handle window sizes
   of up to 2^64 - 1 bytes.  The window MUST NOT be increased above
   2^64 - 1 bytes.

   Data transfer is done with messages of the following type.

      byte      QMUX_MSG_CHANNEL_DATA
      uint64    recipient channel
      string    data

   The maximum amount of data allowed is determined by the maximum
   packet size for the channel, and the current window size, whichever
   is smaller.  The window size is decremented by the amount of data
   sent.  Both parties MAY ignore all extra data sent after the allowed
   window is empty.

   Implementations are expected to have some limit on the transport
   layer packet size.

###  Closing a Channel

   When a party will no longer send more data to a channel, it SHOULD
   send QMUX_MSG_CHANNEL_EOF.

      byte      QMUX_MSG_CHANNEL_EOF
      uint64    recipient channel

   No explicit response is sent to this message.  However, the
   application may send EOF to whatever is at the other end of the
   channel.  Note that the channel remains open after this message, and
   more data may still be sent in the other direction.  This message
   does not consume window space and can be sent even if no window space
   is available.

   When either party wishes to terminate the channel, it sends
   QMUX_MSG_CHANNEL_CLOSE.  Upon receiving this message, a party MUST
   send back an QMUX_MSG_CHANNEL_CLOSE unless it has already sent this
   message for the channel.  The channel is considered closed for a
   party when it has both sent and received QMUX_MSG_CHANNEL_CLOSE, and
   the party may then reuse the channel number.  A party MAY send
   QMUX_MSG_CHANNEL_CLOSE without having sent or received
   QMUX_MSG_CHANNEL_EOF.

      byte      QMUX_MSG_CHANNEL_CLOSE
      uint32    recipient channel

   This message does not consume window space and can be sent even if no
   window space is available.

   It is RECOMMENDED that all data sent before this message be delivered
   to the actual destination, if possible.

## Summary of Message Numbers

   The following is a summary of messages and their associated message
   number byte value.

            QMUX_MSG_CHANNEL_OPEN                    100
            QMUX_MSG_CHANNEL_WINDOW_ADJUST           101
            QMUX_MSG_CHANNEL_DATA                    102
            QMUX_MSG_CHANNEL_EOF                     103
            QMUX_MSG_CHANNEL_CLOSE                   104

## Data Type Representations Used

   byte

      A byte represents an arbitrary 8-bit value (octet).  Fixed length
      data is sometimes represented as an array of bytes, written
      byte[n], where n is the number of bytes in the array.

   uint64

      Represents a 64-bit unsigned integer.  Stored as eight bytes in
      the order of decreasing significance (network byte order).

   string

      Arbitrary length binary string.  Strings are allowed to contain
      arbitrary binary data, including null characters and 8-bit
      characters.  They are stored as a uint64 containing its length
      (number of bytes that follow) and zero (= empty string) or more
      bytes that are the value of the string.  Terminating null
      characters are not used.